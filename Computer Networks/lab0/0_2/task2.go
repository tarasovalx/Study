package main

import (
	"fmt";
	"github.com/mmcdole/gofeed"
	)


type ItemView struct {
	Title string
	Content string
	Source string
}

var urls []string
var feedParser = gofeed.NewParser()

func ParseRSS() {
	for _, url := range urls {
		feedData, _ := feedParser.ParseURL("http://" + url)
		for _, item := range feedData.Items {
			fmt.Println(item.Title)
			fmt.Println(item.Description)
		}
	}
}
func main() {
	urls = []string{"lenta.ru/rss", "news.ap-pa.ru/rss.xml"}
	ParseRSS()
}