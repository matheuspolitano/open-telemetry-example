package main

import (
	"log"
	"os"

	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/matheuspolitano/open-telemetry-example/config"
	"github.com/matheuspolitano/open-telemetry-example/internal/server"
	"github.com/matheuspolitano/open-telemetry-example/pkg/logger"
)

func main() {
	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	appLogger := logger.NewApiLogger(cfg)
	s := server.NewServer(appLogger, *cfg)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
