package main

import (
	"context"
	"lang-bot/internal/app"
	"os"
)

func main() {
	ctx := context.Background()

	app := app.NewApp(ctx, os.Getenv("TELEGRAM_BOT_APIKEY"))
	app.BotUpdate(ctx)
}
