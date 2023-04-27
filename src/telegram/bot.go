package telegram

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

type PinBot struct {
	BotApi    *tgbotapi.BotAPI
	BotConfig tgbotapi.UpdateConfig
}

func (bot *PinBot) Start() {
	bot.BotConfig.Timeout = 60
	updates, _ := bot.BotApi.GetUpdatesChan(bot.BotConfig)

	for update := range updates {
		bot.Do(&update)
	}
}

func (bot *PinBot) Pin(update *tgbotapi.Update) {
	_, err := bot.BotApi.PinChatMessage(tgbotapi.PinChatMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
	})
	if err != nil {
		log.Error(err)
	}
}

func (bot *PinBot) PostChannel(update *tgbotapi.Update) {
	channelId := config.GetInt64("channel_post")
	msg := tgbotapi.NewForward(channelId, update.Message.Chat.ID, update.Message.MessageID)
	_, err := bot.BotApi.Send(msg)
	if err != nil {
		log.Error(err)
	}
}

func (bot *PinBot) Do(update *tgbotapi.Update) {
	if update.Message == nil && update.ChannelPost == nil { // ignore any non-Message Updates
		return
	}

	if update.Message != nil {
		spew.Dump("update message", update.Message)

		if strings.Contains(update.Message.Text, "#закреп") {
			bot.Pin(update)
			bot.PostChannel(update)
		}
	} else {
		spew.Dump("chanel post", update.ChannelPost)
	}
}
