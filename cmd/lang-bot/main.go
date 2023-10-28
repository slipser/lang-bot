package main

import (
	"lang-bot/internal/app"
	"os"
)

func main() {
	app := app.NewApp(os.Getenv("TELEGRAM_BOT_APIKEY"))
	app.BotUpdate()
}
