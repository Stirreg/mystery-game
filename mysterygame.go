package main

import (
	"html/template"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	mysterygameTemplate, _ := template.ParseFiles("templates/mysterygame.html")

	var data = map[string]string{
		"title": "Mystery Game",
	}

	mysterygameTemplate.Execute(writer, data)
}

func main() {
	addr := os.Getenv("APP_ADDR")

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
