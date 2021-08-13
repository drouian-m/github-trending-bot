package scraper

import (
	"fmt"
	"github-trending-bot/models"
	"github.com/gocolly/colly"
	"log"
	"regexp"
	"strings"
)

func GetProjects(language string) ([]models.Project, error) {
	log.Printf("Get %s projects trendings", language)
	baseUrl := "https://github.com"
	url := fmt.Sprintf("%s/trending/%s", baseUrl, language)
	c := colly.NewCollector()

	projects := make([]models.Project, 0)

	c.OnHTML(".Box > div", func(e *colly.HTMLElement) {
		e.ForEach(".Box-row", func(index int, element *colly.HTMLElement) {
			m := regexp.MustCompile(`(\n|\s)+`)
			chars := regexp.MustCompile(`(\n|\s|[a-zA-Z])+`)
			title := m.ReplaceAllString(element.ChildText(".h3"), "")
			description := strings.TrimSpace(element.ChildText(".col-9"))
			url := element.ChildAttr(".h3 > a", "href")
			stars := chars.ReplaceAllString(element.ChildText("div > a"), "")

			project := models.Project{
				Title:       title,
				Description: description,
				URL:         fmt.Sprintf("%s%s", baseUrl, url),
				Stars: stars,
			}
			projects = append(projects, project)
		})
	})

	c.Visit(url)

	return projects, nil
}
