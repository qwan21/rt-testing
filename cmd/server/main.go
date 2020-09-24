package main

import (
	"log"
	"rt/config"
	"rt/internal/server"

	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	server.Start(viper.GetString("server.host"), viper.GetInt("server.port"))
}
