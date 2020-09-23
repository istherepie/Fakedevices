package cli

import (
	"errors"
	"flag"
)

// Config contains the config parameters for the entire application
type Config struct {
	DeviceFile string
}

// CreateConfiguration parses commandline arguments and returns `Config` struct
func CreateConfiguration() (Config, error) {

	// Config
	config := Config{}

	// REQUIRED arg -d <filename>
	// This should point to the `device` file
	file := flag.String("d", "", "Path to a yaml formatted device file.")
	flag.Parse()

	if *file == "" {
		flag.Usage()
		return config, errors.New("MISSING_CLI_ARGS")
	}

	// Set config value
	config.DeviceFile = *file

	return config, nil
}
