package main

import (
	"log"
	"os"
	"strings"

	"customerio/devices/scripts/utils"
)

const (
	SupportedDevicesCsv   = "https://storage.googleapis.com/play_public/supported_devices.csv"
	AndroidDataFile       = "../../src/data/android.json"
	AndroidDataBackupFile = "../../src/data/android.json.bak"
)

type Device struct {
	RetailBranding string `json:"retailBranding"`
	MarketingName  string `json:"marketingName"`
	Device         string `json:"device"`
	Model          string `json:"model"`
}

func main() {
	log.Println("Start fetching android device data from: ", SupportedDevicesCsv)

	// 1. back up current data file
	// 2. download csv from SupportedDevicesCsv
	// 3. write csv -> json to the data file
	// 4. delete backup file

	// 1. back up current data file
	if _, err := utils.CopyFile(AndroidDataFile, AndroidDataBackupFile); err != nil {
		panic(err)
	} else {
		log.Println("Created backup file:", AndroidDataBackupFile)
		defer func() {
			// 4. delete backup file if everything went well
			if err := os.Remove(AndroidDataBackupFile); err != nil {
				panic(err)
			} else {
				log.Println("Deleted backup file:", AndroidDataBackupFile)
			}
		}()
	}

	// 2. download csv from SupportedDevicesCsv
	data, err := utils.ReadCsvFromUrl(SupportedDevicesCsv, ',', true, true)
	if err != nil {
		panic(err)
	}

	log.Println("Fetched", len(data), "devices")

	filteredDevices := filterDevices(data)

	log.Println("Filtered to", len(filteredDevices), "devices based on market share")

	brands := make(map[string]bool)
	for _, row := range filteredDevices {
		brands[row[0]] = true
	}

	log.Println("Found", len(brands), "brands")

	// 3. write csv -> json to the data file
	if err := utils.WriteJsonToFile(filteredDevices, AndroidDataFile, "", "  "); err != nil {
		log.Fatal("Error writing to JSON, check out the *.bak file to restore old contents", err)
	} else {
		log.Println("Wrote", len(filteredDevices), "devices to", AndroidDataFile)
		if fi, err := os.Stat(AndroidDataFile); err == nil {
			log.Printf("File size: %d KB", fi.Size()/1024)
		}
	}
}

func filterDevices(data [][]string) [][]string {
	var filtered [][]string
	for _, row := range data {
		if isRetailsBrandingValid(row[0]) {
			filtered = append(filtered, row)
		}
	}
	return filtered
}

func isRetailsBrandingValid(retailBranding string) bool {
	topManufacturers := []string{
		"samsung",
		"xiaomi",
		"oppo",
		"vivo",
		"huawei",
		"realme",
		"motorola",
		"amazon",
		"infinix",
		"tecno",
		"lge",
		"amlogic",
		"oneplus",
		"google",
		"hmd global",
		"Sony",
		"rockchip",
		"zte",
		"tct (alcatel)",
		"itel",
	}

	for _, manufacturer := range topManufacturers {
		retailBranding = strings.ToLower(retailBranding)
		if retailBranding == manufacturer {
			return true
		}
	}
	return false
}
