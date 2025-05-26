package handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ContainsKeywords(text string, keywords []string) bool{
	text = strings.ToLower(text)
	for _, keyword := range keywords{
		if strings.Contains(text, strings.ToLower(keyword)){
			return true
		}
	}

	return false
}

func ChooseMessages(update tgbotapi.Update, bot *tgbotapi.BotAPI, keywords []string, targetChatID int64){
	
	text := update.Message.Text
	if text == ""{
		text = update.Message.Caption
	}

	if ContainsKeywords(text, keywords){
		msg := tgbotapi.NewMessage(targetChatID, text)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}


}