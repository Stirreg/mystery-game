package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Name  string
	Title string
}

func page(name string) Page {
	data, _ := ioutil.ReadFile("data/pages.json")

	var pages []Page

	err := json.Unmarshal(data, &pages)
	if err != nil {
		fmt.Println(err)
	}

	var page Page

	for _, value := range pages {
		if value.Name == name {
			page = value
		}
	}

	return page
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mysterygameTemplate, _ := template.ParseFiles("templates/mysterygame.html")

	home := page("home")

	mysterygameTemplate.Execute(writer, home)
}

func main() {
	addr := os.Getenv("APP_ADDR")

	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
