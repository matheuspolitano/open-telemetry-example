package main

import (
	"log"
	"os"

	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/matheuspolitano/open-telemetry-example/config"
	"github.com/matheuspolitano/open-telemetry-example/internal/server"
)

func main() {
	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	appLogger := logger.NewApiLogger(cfgFile)
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(appLogger, *cfg)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
