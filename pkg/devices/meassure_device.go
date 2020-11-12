package devices

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Meassure struct {
	Identifier string
	Maker      string
	Value      float32
}

// MeassuringDevice is a device which randomly emits a float value
func MeassuringDevice(device Device, client MQTT.Client, wg *sync.WaitGroup) {

	defer wg.Done()

	payload := Meassure{
		Identifier: device.Name,
		Maker:      device.Maker,
	}

	for {

		payload.Value = rand.Float32()

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
