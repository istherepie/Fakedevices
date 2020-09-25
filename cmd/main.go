package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/istherepie/fakedevices/pkg/app"
	"github.com/istherepie/fakedevices/pkg/devices"
)

func RunApplication() int {

	var wg sync.WaitGroup

	config, err := app.CreateConfiguration()

	if err != nil {
		fmt.Printf("Application could not start - %v\n", err)
		return 1
	}

	contents, err := ioutil.ReadFile(config.DeviceFile)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	importedDevices, err := devices.ImportDevices(contents)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	wg.Add(1)
	client, err := app.CreateClientConnection(config.BrokerAddr(), &wg)

	// Workaround: wait for connection - should be refactored!
	wg.Wait()

	if err != nil {
		fmt.Println(err)
		return 1
	}

	// Start fake devices

	for _, device := range importedDevices {

		var enabled bool = false

		switch device.Type {
		case "meassure":
			wg.Add(1)
			go devices.MeassuringDevice(device, client, &wg)
			enabled = true

		case "switch":
			wg.Add(1)
			go devices.SwitchDevice(device, client, &wg)
			enabled = true
		}

		if enabled {
			fmt.Printf("+ Device: %v enabled!\n", device.Name)
		} else {
			fmt.Printf("- Device: %v is of unknown type and not enabled!\n", device.Name)
		}
	}

	wg.Wait()
	return 0
}
func main() {
	retCode := RunApplication()

	os.Exit(retCode)
}
