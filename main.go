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

	verification := tgbotapi.NewMessage(1714711619, "Bot is up.")
	_, err = bot.Send(verification)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() { // listen to command messages only
			switch update.Message.Command() {
			case "current":
				userInput := update.Message.CommandArguments()

				if userInput == "" {

					_, err = bot.Send(internal.Current(update, "Buenos_Aires"))
					if err != nil {
						log.Println(err)
					}
				} else {

					_, err = bot.Send(internal.Current(update, strings.ReplaceAll(userInput, " ", "-")))
					if err != nil {
						log.Println(err)
					}
				}
			case "ping":

				t := time.Now().Format("02/04/06, 15:04:05")

				msgContent := fmt.Sprintf("I'm alive, server time is %v", t)
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
				err := internal.Handshake(args[0], args[1], args[2])
				if err != nil {
					message := tgbotapi.NewMessage(update.Message.Chat.ID, "error buscando la imagen, seguro no se subio al servidor")
					_, err = bot.Send(message)
					if err != nil {
						log.Println(err)
					}
				}
				path := "/assets/output.jpg"
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
			case "createfile":

				file, err := os.Create("/assets/example.txt")
				if err != nil {
					log.Println(err)
				}
				defer file.Close()

				_, err = file.WriteString("Hello, world!\n")
				if err != nil {
					log.Println(err)
				}

				fmt.Println("File created successfully!")
			case "nextdays":
				userInput := update.Message.CommandArguments()

				if userInput == "" {

					_, err = bot.Send(internal.Forecast(update, "Buenos_Aires"))
					if err != nil {
						log.Println(err)
					}
				} else {

					_, err = bot.Send(internal.Forecast(update, strings.ReplaceAll(userInput, " ", "&20")))
					if err != nil {
						log.Println(err)
					}
				}

			}
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
