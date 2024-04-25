package main

import "fmt"

type Trend struct {
	Title       string
	Description string
	Website     string
	Author      string
}

var urls = []string{"google.com", "youtube.com", "twitter.com"}

func main() {
	trendsChannel := make(chan Trend)

	for _, url := range urls {
		go func(keyword string, website string) {
			newTrend := Trend{Title: "some Topic", Description: "this descripition", Website: url, Author: "this Author"}
			trendsChannel <- newTrend
		}("tech", url)
	}

	trends := []Trend{}
	for range urls {
		trends = append(trends, <-trendsChannel)
	}

	fmt.Println(trends)
}
