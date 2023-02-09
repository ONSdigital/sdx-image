package main

import (
	"sdxImage/pkg/api"
	"sdxImage/pkg/log"
)

func main() {
	log.Init()
	log.Info("Starting sdx-image")
	api.Listen()
}
