package main

import (
	"sdxImage/internal/api"
	"sdxImage/internal/log"
)

func main() {
	log.Init()
	log.Info("Starting sdx-image")
	api.Listen()
}
