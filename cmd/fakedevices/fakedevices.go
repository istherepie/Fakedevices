package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

	for _, device := range importedDevices {
		fmt.Println(device)
	}

	var mqttBrokerAddr string = "tcp://localhost:1883"

	client := mqtt.CreateClientConnection(mqttBrokerAddr)

	msg := "This is my message"
	token := client.Publish("steffen/test/me", 0, false, msg)
	token.Wait()
}
