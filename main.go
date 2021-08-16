package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/evleria/position-client/internal/cmd"
	"github.com/evleria/position-client/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := new(config.Сonfig)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := cmd.NewRootCmd(cfg)
	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
