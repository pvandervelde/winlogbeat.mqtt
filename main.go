package main

import (
	"os"

	_ "github.com/sidleal/mqttout"

	"github.com/elastic/beats/winlogbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
