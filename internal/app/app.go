package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Bot *tgbotapi.BotAPI
}

func NewApp(apikey string) App {
	bot := StartBot(apikey)

	return App{
		Bot: bot,
	}
}

func StartBot(apikey string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func (a App) BotUpdate() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			a.Bot.Send(msg)
		}
	}
}
