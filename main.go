package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	TelegramBotID := flag.String("botid", "", "The Bot Id")
	ChatID := flag.Int64("chatid", -1, "The Chat Id")

	bot, err := tgbotapi.NewBotAPI(*TelegramBotID)

	if err != nil {
		log.Panic(err)
	}

	EventType := os.Getenv("sonarr_eventtype")

	if EventType == "Test" {
		msg := tgbotapi.NewMessage(
			*ChatID,
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
		*ChatID,
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
