package scraper

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func ScrapeTitles(url string, queryClass string) []string {
	fmt.Println("Started scraping")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run tasks
	// Replace "http://example.com" with the URL you want to scrape
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		// Navigate to site
		chromedp.Navigate(url),
		// Wait for an element that is rendered by JavaScript
		chromedp.WaitVisible(queryClass, chromedp.ByQueryAll),
		// Retrieve the content of the element
		chromedp.Nodes(queryClass, &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the extracted content
	fmt.Println("Printing Response Nodes")

	titles := []string{}
	for _, node := range nodes {
		titles = append(titles, node.Children[0].NodeValue)
	}

	return titles
}
