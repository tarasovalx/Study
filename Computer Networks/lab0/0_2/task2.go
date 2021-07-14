package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
)

var urls = []string{"lenta.ru/rss", "news.ap-pa.ru/rss.xml"}
var feedParser = gofeed.NewParser()

func ParseRSS(url string) {
	feedData, err := feedParser.ParseURL("http://" + url)
	if err != nil { log.Fatal(err) }
	for _, item := range feedData.Items {
		fmt.Println(item.Title)
		fmt.Println(item.Description)
	}
}

func main() {
	for _, url := range urls {
		ParseRSS(url)
	}
}