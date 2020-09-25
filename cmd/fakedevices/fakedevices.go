package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/istherepie/fakedevices/pkg/cli"
	"github.com/istherepie/fakedevices/pkg/devices"
	"github.com/istherepie/fakedevices/pkg/mqtt"
)

func main() {
	config, err := cli.CreateConfiguration()

	fmt.Println(config)

	if err != nil {
		fmt.Printf("Application could not start - %v\n", err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(config.DeviceFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	importedDevices, err := devices.ImportDevices(contents)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := mqtt.CreateClientConnection(config.BrokerAddr())

	// Start fake devices

	var wg sync.WaitGroup

	for _, device := range importedDevices {
		fmt.Printf("# Enabling device: %v\n", device.Name)

		switch device.Type {
		case "meassure":
			wg.Add(1)
			go devices.MeassuringDevice(device, client, &wg)
		case "switch":
			wg.Add(1)
			go devices.SwitchDevice(device, client, &wg)
		default:
			fmt.Printf("Device: %v is of unknown type (%v)", device.Name, device.Type)
		}
	}

	wg.Wait()
}
