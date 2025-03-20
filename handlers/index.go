package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
}

func Index(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := PageData{
		Title: "Index",
		Body:  "Some body"}
	err := tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
