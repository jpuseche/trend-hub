package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ScrapeData() {
	fmt.Println("Started scraping")

	c := colly.NewCollector()

	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

}
