package main

import (
	"fmt"
	"os"

	"github.com/istherepie/fakedevices/pkg/cli"
)

func main() {
	config, err := cli.CreateConfiguration()

	if err != nil {
		fmt.Printf("Application could not start - %v\n", err)
		os.Exit(1)
	}
}
