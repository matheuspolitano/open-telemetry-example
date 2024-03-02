package main

import (
	"log"

	"github.com/matheuspolitano/open-telemetry-example/config"
)

func main() {
	cfgFile, err := config.LoadConfig("./config/local-conf")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
}
