package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	TelegramBotID := os.Getenv("BOT-ID")
	ChatID, err := strconv.ParseInt(os.Getenv("CHAT-ID"), 10, 64)

	if err != nil {
		log.Fatal("Invalid Chat ID provided")
	}

	bot, err := tgbotapi.NewBotAPI(TelegramBotID)

	if err != nil {
		log.Panic(err)
	}

	EventType := os.Getenv("sonarr_eventtype")

	if EventType == "Test" {
		msg := tgbotapi.NewMessage(
			ChatID,
			"Test Successful!",
		)
		bot.Send(msg)
		return
	}

	SeriesTitle := os.Getenv("sonarr_series_title")
	SeasonNumber := os.Getenv("sonarr_episodefile_seasonnumber")
	EpisodeNumbers := os.Getenv("sonarr_episodefile_episodenumbers")
	ImdbID := os.Getenv("sonarr_series_imdbid")
	Quality := os.Getenv("sonarr_episodefile_quality")

	msg := tgbotapi.NewMessage(
		ChatID,
		fmt.Sprintf(
			"Episode Downloaded\n %s - S%sE%s (%s) - https://www.imdb.com/title/%s/",
			SeriesTitle,
			SeasonNumber,
			EpisodeNumbers,
			Quality,
			ImdbID,
		),
	)

	bot.Send(msg)
}
