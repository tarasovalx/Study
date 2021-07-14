package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var PORT = ":9034"
var urls = map[string]string {
	"/mail.ru/" : "http://mail.ru",
	"/ya.ru/" : "http://ya.ru",
	"/lenta.ru" : "http://lenta.ru",
}

type Item struct {
	Name string
	Path string
	url string
}

type ViewData struct {
	Items []Item
}

var t *template.Template
var data *ViewData

func HomeRouterHandler(w http.ResponseWriter, r *http.Request){
	t.Execute(w, data)
}

func PageRouteHandler(w http.ResponseWriter, r *http.Request){
	res, err := http.Get("https:/" + r.URL.Path)
	if err != nil { log.Fatal(err) }
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { log.Fatal(err) }
	_, err = fmt.Fprintf(w, string(body))
	if err != nil { log.Fatal(err) }
}	

func Init() {
	t, _ = template.ParseFiles("index.html")
	data = &ViewData{
		Items : []Item{
			{Name: "mail", Path: "/mail.ru/", url: "http://mail.ru"},
			{Name: "yandex", Path: "/ya.ru/", url: "http://ya.ru"},
			{Name: "lenta", Path: "/lenta.ru/", url: "http://lenta.ru"}}}
}

func main() {
	Init()

	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/mail.ru/", PageRouteHandler)
	http.HandleFunc("/ya.ru/", PageRouteHandler)
	http.HandleFunc("/lenta.ru/", PageRouteHandler)
	err := http.ListenAndServe(PORT, nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}