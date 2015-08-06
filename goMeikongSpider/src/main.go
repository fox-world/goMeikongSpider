package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"log"
	"net/http"
	"service"
)

func spider(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/spider.html")
	t.Execute(w, nil)
}

func main() {

	var url = "http://www.moko.cc/channels/post/23/1.html"
	var num = 0
	http.Handle("/", websocket.Handler(service.EchoServer(url, num)))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/spider", spider)
	if err := http.ListenAndServe(":8004", nil); err != nil {
		log.Fatal("ListentAndServe:", err)
	}
}
