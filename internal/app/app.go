package app

import (
	"lang-bot/internal/bot"
	translationService "lang-bot/internal/translation/service"
	"lang-bot/internal/word"
	"lang-bot/internal/word/repository/localcache"
	wordService "lang-bot/internal/word/service"
)

type App struct {
	Bot         bot.Bot
	WordService word.Service
}

func NewApp(apikey string) App {
	translationService := translationService.NewTranslationService()

	wordRepository := localcache.NewWordLocalStorage()
	wordService := wordService.NewWordService(wordRepository, translationService)

	bot := bot.NewBot(apikey, wordService)

	return App{
		Bot:         bot,
		WordService: wordService,
	}
}
