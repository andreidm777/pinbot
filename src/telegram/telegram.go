package telegram

import (
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	config.SetDefault("bot_secret", "")
	config.SetDefault("channel_post", 0)
}

func Run() {
	b, err := tgbotapi.NewBotAPI(config.GetString("bot_secret"))
	if err != nil {
		log.Panic(err)
	}

	log.Debugf("Authorized on account %s", b.Self.UserName)

	bot := &PinBot{
		BotApi:    b,
		BotConfig: tgbotapi.NewUpdate(0),
	}

	bot.Start()
}
