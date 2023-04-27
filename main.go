package main

import (
	"flag"
	"pinbot/src/telegram"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

var configFileName = flag.String("config", "/usr/local/etc/pinbot.conf", "config file")

func main() {
	flag.Parse()
	config.AllowEmptyEnv(true)
	config.SetConfigType("yaml")
	config.SetConfigFile(*configFileName)

	if err := config.ReadInConfig(); err != nil {
		log.Error(err)
	}

	config.WatchConfig()

	telegram.Run()
}
