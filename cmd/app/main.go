package main

import (
	"log"

	"github.com/polluxdev/financing-system/config"
	"github.com/polluxdev/financing-system/internal/app"
)

func main() {
	// load config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// run app
	app.Run(config)
}
