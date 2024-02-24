package main

import (
	"gofinance/config"
	"gofinance/web"

	"log"

	"github.com/spf13/pflag"
)

func loadConfig() *config.AppCfg {
	filepath := pflag.StringP("config", "c", "", "configuration filepath (default: None)")
	pflag.Parse()
	// ________________________________________________________________________
	// Load config
	cfg, err := config.NewAppCfg(*filepath)
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
		panic("Could not load config. Shutting down")
	}
	return cfg
}

func main() {
	cfg := loadConfig()
	web := web.Create(cfg)
	web.Run()

	// handle graceful shutdown
	// web.Shutdown();
}
