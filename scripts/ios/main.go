package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map"
)

type Device struct {
	Identifier string `json:"identifier"`
	Generation string `json:"generation"`
}

func main() {
	// Fetch devices raw data from https://github.com/pluwen/apple-device-model-list
	responseAsLines, err := readDevicesRawDataFromGithubAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer devicesRawResponse.Body.Close()

	// We are going to be processing the ReadMe file line by line, those variables will help us keep state

	// The read me file has different sections for different devices (iPhone, iPad...etc)
	// This keeps track of the current section being processed
	var currentSection string
	// Boolean variable to indicate if we are currently processing lines in a code block
	isCodeBlock := false

	// A map to keep the overall resultMap
	// Key: string -> Section name (iPhone, iPad...etc)
	// Value : Map
	//        - Key: string -> Device Generation
	//        - Value: []string -> Device identifiers for a single generation
	resultMap := orderedmap.New()

	for _, line := range responseAsLines {

		if strings.HasPrefix(line, "## ") {
			// New section detected
			currentSection = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			resultMap.Set(currentSection, orderedmap.New())
			continue
		}

		if strings.HasPrefix(line, "```") {
			// Code block start or end detected
			isCodeBlock = !isCodeBlock
			continue
		}

		if isCodeBlock && currentSection != "" {
			blockResult, _ := resultMap.Get(currentSection)
			typedBlockResult := blockResult.(*orderedmap.OrderedMap)

			// Parse a single code block line
			// Example: `"iPhone3,1", "iPhone3,2", "iPhone3,3":             iPhone 4`
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				generation := strings.TrimSpace(parts[1])
				identifiers := extractIdentifiers(strings.TrimSpace(parts[0]))
				typedBlockResult.Set(generation, identifiers)
			}

			resultMap.Set(currentSection, typedBlockResult)
		}
	}

	// Flatten result into list of OutputItems
	devices := make([]Device, 0)
	devices = append(devices, addSectionDevicesToOutputList("Apple TV", resultMap)...)
	devices = append(devices, addSectionDevicesToOutputList("Apple Watch", resultMap)...)
	devices = append(devices, addSectionDevicesToOutputList("iPad", resultMap)...)
	devices = append(devices, addSectionDevicesToOutputList("iPhone", resultMap)...)

	// Write result to JSON file
	outputFile, err := os.Create("../../src/data/ios.json")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(devices); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Data successfully written to:", outputFile.Name())
}

func readDevicesRawDataFromGithubAsLines() ([]string, error) {
	url := "https://raw.githubusercontent.com/pluwen/apple-device-model-list/main/README.md"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("Loaded devices raw data successfully!")

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return lines, nil
}

// This function accepts input like that: `"iPhone3,1", "iPhone3,2", "iPhone3,3"`
// It's tricky to do a simple split because the values we are interested in are separated by ','
// But they also have ',' in their value
// Expected output for above example is: [iPhone3,1, iPhone3,2, iPhone3,3]
func extractIdentifiers(input string) []string {
	result := make([]string, 0)
	isProcessingWord := false
	wordUnderConstruction := ""

	// The order of checks in this loop is important
	for _, char := range input {
		if char == '"' {
			// This could be either start or end of a word
			if isProcessingWord {
				// End of word detected
				result = append(result, wordUnderConstruction)
				wordUnderConstruction = ""
			}
			isProcessingWord = !isProcessingWord
			continue
		}

		if isProcessingWord {
			wordUnderConstruction += string(char)
			continue
		}

		// Do this check last to avoid ignoring commands and spaces within an identifier value
		if char == ',' || char == ' ' {
			continue
		}
	}

	return result
}

func addSectionDevicesToOutputList(sectionName string, data *orderedmap.OrderedMap) []Device {
	devices := make([]Device, 0)

	sectionMap, _ := data.Get(sectionName)
	if innerMap, ok := sectionMap.(*orderedmap.OrderedMap); ok {
		for innerPair := innerMap.Oldest(); innerPair != nil; innerPair = innerPair.Next() {
			for _, identifier := range innerPair.Value.([]string) {
				devices = append(devices, Device{Identifier: identifier, Generation: innerPair.Key.(string)})
			}
		}
	}

	return devices
}
