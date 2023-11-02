package app

import (
	"lang-bot/internal/bot"
	"lang-bot/internal/word"
	"lang-bot/internal/word/repository/localcache"
	"lang-bot/internal/word/service"
)

type App struct {
	Bot         bot.Bot
	WordService word.Service
}

func NewApp(apikey string) App {
	wordRepository := localcache.NewWordLocalStorage()
	wordService := service.NewWordService(wordRepository)

	bot := bot.NewBot(apikey, wordService)

	return App{
		Bot:         bot,
		WordService: wordService,
	}
}
