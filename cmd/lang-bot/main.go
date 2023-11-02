package main

import (
	"context"
	"lang-bot/internal/app"
	"os"
)

func main() {
	ctx := context.Background()

	app := app.NewApp(os.Getenv("TELEGRAM_BOT_APIKEY"))
	app.Bot.BotUpdate(ctx)
}
