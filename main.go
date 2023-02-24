package main

import (
	"fmt"
	"github.com/faculerena/hourlyWeatherBot/internal"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
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
			case "handshake":
				userInput := update.Message.CommandArguments()
				args := strings.Split(userInput, " - ")

				if len(args) != 3 {
					message := tgbotapi.NewMessage(update.Message.Chat.ID, "pone bien las cosas, pusiste "+userInput)
					_, err = bot.Send(message)
					if err != nil {
						log.Println(err)
					}

				}
				internal.Handshake(args[0], args[1], args[2])
				path := "./output/output.jpg"
				file, err := os.Open(path)
				if err != nil {
					log.Println(err)
				}
				outputName := "handshake-" + userInput + ".jpg"

				message := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, tgbotapi.FileReader{
					Name:   outputName,
					Reader: file,
					Size:   -1,
				})
				_, err = bot.Send(message)
				if err != nil {
					log.Println(err)
				}
				err = file.Close()
				if err != nil {
					log.Println(err)
				}
			}

		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
