package main

import (
	"html/template"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	mysterygameTemplate, _ := template.ParseFiles("templates/mysterygame.html")

	var data = map[string]string{
		"title": "Mystery Game",
	}

	mysterygameTemplate.Execute(writer, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
