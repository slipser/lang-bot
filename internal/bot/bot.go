package bot

import (
	"context"
	"lang-bot/internal/word"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type wordService interface {
	AddWord(ctx context.Context, text string) (word.Word, error)
	GetWords(ctx context.Context) ([]word.Word, error)
}

type Bot struct {
	bot         *tgbotapi.BotAPI
	wordService wordService
}

func NewBot(apikey string, wordService wordService) Bot {
	bot, err := startBot(apikey)
	if err != nil {
		log.Printf("Error authorize, %v", err)
		return Bot{}
	}

	return Bot{
		bot:         bot,
		wordService: wordService,
	}
}

func startBot(apikey string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}

func (b Bot) BotUpdate(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	isWaitingWord := false

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			if isWaitingWord {
				isWaitingWord = false
				_, err := b.wordService.AddWord(ctx, update.Message.Text)
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error"+err.Error())
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "add!")
				}
			}

			switch update.Message.Text {
			case "/add":
				isWaitingWord = true
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Enter word")
			case "/get":
				words, err := b.wordService.GetWords(ctx)
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error"+err.Error())
					continue
				}

				allWords := "words: "

				for _, word := range words {
					allWords += word.Text + " : " + word.Translation + " "
				}

				msg = tgbotapi.NewMessage(update.Message.Chat.ID, allWords)
			}

			msg.ReplyToMessageID = update.Message.MessageID

			b.bot.Send(msg)
		}
	}
}
