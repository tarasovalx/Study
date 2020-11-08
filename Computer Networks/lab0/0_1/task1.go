package main

import (
	"fmt";
	"net/http";
	"io/ioutil";
	"html/template";
	"log";
	)

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

var port string
var t *template.Template

var data *ViewData

func HomeRouterHandler(w http.ResponseWriter, r *http.Request){
	t.Execute(w, data)
}

func PageRouteHandler(w http.ResponseWriter, r *http.Request){
	res, err := http.Get("https:/" + r.URL.Path)
	if err != nil { log.Fatal(err) }
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, string(body))
}	

func Init() {
	port = ":9034"
	t, _ = template.ParseFiles("index.html") 

	data = &ViewData{
		Items : []Item{
			Item{Name : "mail", Path : "/mail.ru/", url : "http://mail.ru"},
			Item{Name : "yandex", Path : "/ya.ru/", url : "http://ya.ru"},
			Item{Name : "lenta", Path : "/lenta.ru/", url : "http://lenta.ru"},}}
}

func main() {
	Init()
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/mail.ru/", PageRouteHandler)
	http.HandleFunc("/ya.ru/", PageRouteHandler)
	http.HandleFunc("/lenta.ru/", PageRouteHandler)
	err := http.ListenAndServe(port, nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}