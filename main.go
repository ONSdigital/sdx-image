package main

import (
	"os"
	"sdxImage/internal/api"
	"sdxImage/internal/log"
)

func main() {
	log.Init()
	version := readVersion()
	log.Info("Starting sdx-image version: " + version)
	api.Listen()
}

func readVersion() string {
	data, err := os.ReadFile(".version")
	if err != nil {
		log.Error("Failed to read .version file: %v", err)
	}

	return string(data)
}
