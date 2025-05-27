package main

import (
	"log"
	"os"
	"telegram-bot-go/clients"
	// "telegram-bot-go/config"
	"telegram-bot-go/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)


func main(){
	// _ = config.Config("TELEGRAM_APITOKEN")
	_ = os.Getenv("TELEGRAM_APITOKEN")

	bot := clients.Init()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Telegram bot is running")
	})

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Fatal(app.Listen(":" + port))
	}()


	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	keywords := []string{"golang", "go", "junior"}
	targetChatID := int64(-1002471834156)

	for update := range updates {
		if update.Message != nil {
			handlers.ChooseMessages(update, bot, keywords, targetChatID)
		}
	}
}