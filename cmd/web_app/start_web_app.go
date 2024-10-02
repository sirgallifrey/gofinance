package main

import (
	"requirementor/config"
	"requirementor/web"
)

func main() {
	cfg := config.LoadConfig()
	web := web.Create(cfg)
	web.Run()

	// handle graceful shutdown
	// web.Shutdown();
}
