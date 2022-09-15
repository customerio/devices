package main

import (
	"customerio/devices/scripts/utils"
	"log"
	"os"
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

	devices := make([]Device, 0, len(data))
	for _, row := range data {
		devices = append(devices, Device{
			RetailBranding: row[0],
			MarketingName:  row[1],
			Device:         row[2],
			Model:          row[3],
		})
	}

	// 3. write csv -> json to the data file
	if err := utils.WriteJsonToFile(devices, AndroidDataFile); err != nil {
		log.Fatal("Error writing to JSON, check out the *.bak file to restore old contents", err)
	} else {
		log.Println("Wrote", len(devices), "devices to", AndroidDataFile)
	}
}
