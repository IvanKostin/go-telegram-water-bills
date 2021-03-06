package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		i, err := strconv.Atoi(msg.Text)
		if err != nil {
			// handle error
			msg.Text = "Enter correct number"
		} else {
			msg.Text = fmt.Sprintf("Water in rubles for month - %d", i*61)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
