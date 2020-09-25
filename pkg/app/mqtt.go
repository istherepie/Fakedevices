package app

import (
	"fmt"
	"sync"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// We present ourself as the app
var clientID string = "fakedevicesApp"

// CreateClientConnection connects to a MQTT (e.g. Mosquitto) service and returns a client struct
// TODO: Rather than using a sync.WaitGroup, a `OnConnect` handler should be passed in... soon...
func CreateClientConnection(brokerAddr string, wg *sync.WaitGroup) (MQTT.Client, error) {

	// Configure
	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerAddr)
	opts.SetClientID(clientID)

	// On connect handler
	opts.OnConnect = func(c MQTT.Client) {
		fmt.Println("Connected to the service")
		wg.Done()
	}

	// Client
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		wg.Done()
		return client, token.Error()
	}

	return client, nil
}
