package main

import (
	"github.com/CodeBitsOrg/dailystats/app"
	"github.com/CodeBitsOrg/dailystats/stats"
	"github.com/CodeBitsOrg/dailystats/telegram"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file was provided")
	}

	portStr := os.Getenv("PORT")

	statsClient := stats.NewStatsClient(http.DefaultClient)

	b, err := bot.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Could not instantiate a new Telegram Bot instance: %s", err)
	}

	tClient := telegram.New(b)

	router := app.Router(app.NewHandler(statsClient, tClient))

	http.ListenAndServe(":"+portStr, router)
}
