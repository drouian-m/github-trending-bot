package main

import (
	"fmt"
	"github-trending-bot/models"
	"github-trending-bot/scraper"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func main() {
	subscriptions := []models.Subscriptions{
		{
			ID:        "1",
			Channel:   1,
			Languages: []string{"js", "go"},
		},
		{
			ID:        "2",
			Channel:   2,
			Languages: []string{"html", "css"},
		},
	}

	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		panic(err)
	}

	//bot.Debug = true

	for _, sub := range subscriptions {
		for _, language := range sub.Languages {
			projects, _ := scraper.GetProjects(language)
			//fmt.Println(projects)

			message := fmt.Sprintf("<strong>Daily Github %s trends</strong>\n\n", strings.ToUpper(language))
			for _, project := range projects {
				message += fmt.Sprintf("<a href=\"%s\">%s  (%s ⭐)</a>\n", project.URL, project.Title, project.Stars)
				if len(project.Description) > 0 {
					message += fmt.Sprintf("<pre>%s</pre>\n", project.Description)
				}
				message += "\n"
			}

			msg := tgbotapi.NewMessage(sub.Channel, message)
			msg.ParseMode = "html"
			msg.DisableWebPagePreview = true

			if _, err := bot.Send(msg); err != nil {
				// Note that panics are a bad way to handle errors. Telegram can
				// have service outages or network errors, you should retry sending
				// messages or more gracefully handle failures.
				panic(err)
			}
			//fmt.Println(message)
		}
	}
}
