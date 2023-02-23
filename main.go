package main

import (
	"fmt"
	"github.com/faculerena/hourlyWeatherBot/internal"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"time"
)

func main() {

	telegramKey := os.Getenv("TELEGRAM_KEY")

	bot, err := tgbotapi.NewBotAPI(telegramKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() { // listen to command messages only
			switch update.Message.Command() {
			case "current":

				_, err = bot.Send(internal.Current(update))
				if err != nil {
					log.Println(err)
				}
			case "ping":

				msgContent := fmt.Sprintf("I'm alive, server time is %v", time.Now())
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgContent)

				_, err = bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}

		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
