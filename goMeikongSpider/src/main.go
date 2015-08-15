package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"log"
	"models"
	"net/http"
	"service"
	"strings"
)

func spider(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/spider.html")
	t.Execute(w, nil)
}

func list(dburi string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.New("list.html").Funcs(funcMap).ParseFiles("template/model/list.html")
		models := service.QueryPage(dburi)
		t.Execute(w, models)
	}
}

var funcMap = template.FuncMap{
	"parseImage": parseImage,
}

func parseImage(str string) string {
	return str[strings.LastIndex(str, "/")+1:]
}

func main() {

	var url = "http://www.moko.cc/channels/post/23/1.html"
	var num = 0

	var config = new(models.Config)
	config.ReadConfig()

	http.Handle("/", websocket.Handler(service.EchoServer(url, num, config.DBuri)))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/模特/", http.StripPrefix("/模特/", http.FileServer(http.Dir("模特"))))
	http.HandleFunc("/spider", spider)
	http.HandleFunc("/list", list(config.DBuri))
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatal("ListentAndServe:", err)
	}
}
