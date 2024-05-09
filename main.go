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

var urls = []string{"google.com", "youtube.com", "twitter.com"}

func home(res http.ResponseWriter, req *http.Request) {
	scraper.ScrapeData()

	startTime := time.Now()
	fmt.Println("Getting Trends")

	trendsChannel := make(chan Trend)

	for _, url := range urls {
		go func(keyword string, website string) {
			time.Sleep(2 * time.Second)

			newTrend := Trend{Title: "some Topic", Description: "this descripition", Website: url, Author: "this Author"}
			trendsChannel <- newTrend
		}("tech", url)
	}

	trends := []Trend{}
	for range urls {
		trends = append(trends, <-trendsChannel)
	}
	endTime := time.Now()

	fmt.Printf("Got Trends sucessfully\n")
	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))
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
