package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"trend-hub/scraper"
)

type Trend struct {
	Title       string
	Description string
	Website     string
	Author      string
}

// This type defines the scraping options to be used per website
type WebScrapeConfig struct {
	URL        string
	QueryClass string
}

// This saves the group of website configs to use on the web scraping
var wscGroup = []WebScrapeConfig{
	{
		URL:        "https://www.tiktok.com/discover/tech-trends",
		QueryClass: ".css-1anth1x-DivVideoDescription",
	},
}

func home(res http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	fmt.Println("Getting Trends")

	trendsChannel := make(chan []Trend)

	// var wg sync.WaitGroup
	for _, wsc := range wscGroup {
		// wg.Add(1)

		go func(keyword string, website string) {
			titles := scraper.ScrapeTitles(wsc.URL, wsc.QueryClass)
			trends := []Trend{}
			for _, title := range titles {
				newTrend := Trend{Title: title, Description: "this descripition", Website: wsc.URL, Author: "this Author"}
				trends = append(trends, newTrend)
			}
			trendsChannel <- trends
		}("tech", wsc.URL)
	}

	// wg.Wait()
	trends := []Trend{}
	for range wscGroup {
		trends = append(trends, <-trendsChannel...)
	}
	endTime := time.Now()

	fmt.Printf("Got Trends sucessfully\n")
	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))
	fmt.Printf("Trends Amount: %v\n", len(trends))
	fmt.Fprint(res, trends)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
