package main

import (
	"fmt"
	"log"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	err := adapter.Enable()
	if err != nil {
		log.Fatal("Unable to enable BLE stack", err)
	}

	// Start scanning.
	fmt.Println("scanning...")
	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		fmt.Println("found device:", device.Address.String(), device.RSSI, device.LocalName())
	})
	if err != nil {
		log.Fatal("Unable to start scan", err)
	}
}
