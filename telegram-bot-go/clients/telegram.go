package clients

import (
	"log"
	"telegram-bot-go/config"

	tbgotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tbgotapi.BotAPI{
	bot, err := tbgotapi.NewBotAPI(config.Config("TELEGRAM_APITOKEN"))
	if err != nil{
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}