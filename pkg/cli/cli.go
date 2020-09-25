package cli

import (
	"errors"
	"flag"
	"fmt"
)

// Config contains the config parameters for the entire application
type Config struct {
	DeviceFile string
	Hostname   string
	Port       int
}

func (c Config) BrokerAddr() string {
	return fmt.Sprintf("tcp://%v:%d", c.Hostname, c.Port)
}

// CreateConfiguration parses commandline arguments and returns `Config` struct
func CreateConfiguration() (Config, error) {

	// Config
	config := Config{}

	// REQUIRED arg -d <filename>
	// This should point to the `device` file
	file := flag.String("d", "", "Path to a yaml formatted device file.")

	// IP addr or hostname of the MQTT broker
	brokerHostname := flag.String("h", "localhost", "Hostname/IP of the MQTT broker service.")

	// IP addr or hostname of the MQTT broker
	brokerPort := flag.Int("p", 1883, "Port of the MQTT broker service.")

	flag.Parse()

	if *file == "" {
		flag.Usage()
		return config, errors.New("MISSING_CLI_ARGS")
	}

	// Set config value
	config.DeviceFile = *file
	config.Hostname = *brokerHostname
	config.Port = *brokerPort

	return config, nil
}
