package main

import (
	"log"

	"github.com/matheuspolitano/open-telemetry-example/config"
)

func main() {
	vconf, err := config.LoadConfig("./config/local-conf")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(vconf)
	if err != nil {
		log.Fatal(err)
	}
}
