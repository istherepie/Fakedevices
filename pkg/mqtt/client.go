package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// We present ourself as the app
var clientID string = "fakedevicesApp"

// CreateClientConnection connects to a MQTT (e.g. Mosquitto) service and returns a client struct
func CreateClientConnection(brokerAddr string) MQTT.Client {

	// Configure
	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerAddr)
	opts.SetClientID(clientID)

	// On connect handler
	opts.OnConnect = func(c MQTT.Client) {
		fmt.Println("Connected to the service")
	}

	// Client
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
