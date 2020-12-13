package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jpoz/switchboard/pkg/conductor"
	"github.com/jpoz/switchboard/pkg/config"
)

func main() {
	configFilePath := os.Args[1]
	if configFilePath == "" {
		panic("missing config")
	}

	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("Could not read %s", configFilePath))
	}

	config, err := config.Parse(buf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse %s: %s", configFilePath, err))
	}

	con := conductor.New(config)
	con.Start()
}
