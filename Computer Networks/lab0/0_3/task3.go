package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type ContentItem struct {
	Items []*NewsItem
	Urls []string
	Names []string
}

type NewsItem struct {
	Title string
	Content template.HTML
	Source string
	Date time.Time
}

var URLS = []string{"lenta.ru/rss", "news.ap-pa.ru/rss.xml"}
var NAMES = []string{"lenta", "news"}
var PORT = ":9034"
var feedParser = gofeed.NewParser()


func MapItem(item *gofeed.Item, Source string) *NewsItem {
	return &NewsItem{Title : item.Title,
					Content : template.HTML(item.Description),
		            Source : Source,
		            Date : *item.PublishedParsed}
}	

func ParseRSS(url string) *[]*NewsItem {
	feedData, err := feedParser.ParseURL("http://" + url)
	if err!= nil { log.Fatal(err) }
	fmt.Println("Parsed: " + url)

	items := make([]*NewsItem, len(feedData.Items))
	for i:= 0; i < len(feedData.Items); i++{
		items[i] = MapItem(feedData.Items[i], feedData.Title)
	}

	return &items
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	allItems := make([]*NewsItem, 0)

	for _, url := range URLS {
		items := ParseRSS(url)
		allItems = append(allItems, *items...)
	}

	sort.SliceStable(allItems, func(i, j int) bool {
		return allItems[i].Date.Unix() > allItems[j].Date.Unix()
	})

	data := ContentItem{Items : allItems, Urls : URLS, Names : NAMES}
	t, _ := template.ParseFiles("./index.html")

	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}


func PageHandler(url string) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		items := ParseRSS(url)
		data := ContentItem{Items : *items, Urls : URLS, Names : NAMES}
		t, _ := template.ParseFiles("./index.html")
		if err := t.Execute(w, data); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)

	for i := 0; i < len(URLS); i++{
		http.HandleFunc("/" + NAMES[i], PageHandler(URLS[i]))
	}

	if err := http.ListenAndServe(PORT, nil); err != nil{
		log.Fatal(err)
	}
}