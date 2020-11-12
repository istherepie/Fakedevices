package devices

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Switch describes the telemetry of the device
type Switch struct {
	Identifier string
	Maker      string
	State      string // State should be `on` or `off`
}

// SwitchDevice is a device which randomly emits an `on/off` state
func SwitchDevice(device Device, client MQTT.Client, wg *sync.WaitGroup) {

	defer wg.Done()

	payload := Switch{
		Identifier: device.Name,
		Maker:      device.Maker,
	}

	for {

		chance := rand.Intn(10)

		if chance < 5 {
			payload.State = "on"
		} else {
			payload.State = "off"
		}

		msg, err := json.Marshal(payload)

		if err != nil {
			fmt.Println(err)
			continue
		}

		token := client.Publish(device.Topic, 0, false, msg)
		token.Wait()

		time.Sleep(time.Duration(device.Interval) * time.Millisecond)
	}
}
