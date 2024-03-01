package main

import (
	"fmt"
	"log"

	"github.com/matheuspolitano/open-telemetry-example/config"
)

func main() {
	vconf, err := config.LoadConfig("./config/local-conf")
	if err != nil {
		log.Fatal(err)
	}
	var conf config.Config
	vconf.Unmarshal(&conf)
	fmt.Println(conf)
}
