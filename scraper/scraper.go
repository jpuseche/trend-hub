package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeData() {
	fmt.Println("Started scraping")

	c := colly.NewCollector()

	c.OnHTML(".body_bold.color-titles", func(e *colly.HTMLElement) {
		fmt.Println("---------------------------------")
		fmt.Println(e.Text)
	})

	scrapingUrl := "https://brightdata.com/"
	err := c.Visit(scrapingUrl)
	if err != nil {
		fmt.Printf("Can't get url for scraping: %v\n", scrapingUrl)
	}

}
