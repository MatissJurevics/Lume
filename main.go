package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
	"Lume/core/config"
)

const protocol = "https"
const host = "openapi.api.govee.com"

// API Response structures
type DeviceResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []Device `json:"data"`
}

type Device struct {
	SKU          string       `json:"sku"`
	Device       string       `json:"device"`
	DeviceName   string       `json:"deviceName"`
	Type         string       `json:"type"`
	Capabilities []Capability `json:"capabilities"`
}

type Capability struct {
	Type       string                 `json:"type"`
	Instance   string                 `json:"instance"`
	Parameters map[string]interface{} `json:"parameters"`
}



func listDevices() {
	url := protocol + "://" + host + "/router/api/v1/user/devices"
	method := "GET"
	
	apiKey := config.GetApiKey()

	if apiKey == "" {
		fmt.Println("No API key found in config. Please set it in the config file.")
		return
	}
	
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Govee-API-Key", apiKey)
    


	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON response
	var deviceResponse DeviceResponse
	err = json.Unmarshal(body, &deviceResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	for i, device := range deviceResponse.Data {
		fmt.Printf("%d. %s\n", i+1, device.DeviceName)
		fmt.Printf("   Type: %s\n", device.Type)
		fmt.Println()
	}
	
}

func printHelp() {
	fmt.Println("Usage: lume <command>")
	fmt.Println("Commands:")
	fmt.Println("  list - List all devices")
	fmt.Println("  help - Show this help message")
	fmt.Println("  set - Set a value in the config")
	fmt.Println("  config - Print the config")
}


func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "list":
		listDevices()
	case "help":
		printHelp()
	case "set":
		config.SetValue(os.Args[2], os.Args[3])
	case "config":
		config.PrintConfig()
	default:
		fmt.Println("Unknown command. Use 'lume help' for available commands.")
		os.Exit(1)
	}
}