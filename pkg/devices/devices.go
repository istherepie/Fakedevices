package devices

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

// Device is the schema of imported devices
type Device struct {
	Name     string
	Topic    string `yaml:"topic"`
	Maker    string `yaml:"maker"`
	Type     string `yaml:"type"`
	Interval int    `yaml:"interval"`
}

// Valid checks that device fields are not empty
func (d Device) Valid() bool {
	return d.Name != "" && d.Topic != "" && d.Maker != "" && d.Type != "" && d.Interval != 0
}

// ImportDevices imports user devices defined in the yaml device file
func ImportDevices(content []byte) ([]Device, error) {

	var validDevices []Device

	importedDevices := map[string]Device{}

	err := yaml.Unmarshal(content, importedDevices)

	if err != nil {
		errorMsg := fmt.Sprintf("Error unmarshalling yaml content - %v\n", err)
		return validDevices, errors.New(errorMsg)
	}

	for name, device := range importedDevices {

		// Set the name of the device
		device.Name = name

		// Append if all fields are set
		if device.Valid() {
			validDevices = append(validDevices, device)
		}
	}

	return validDevices, nil
}
