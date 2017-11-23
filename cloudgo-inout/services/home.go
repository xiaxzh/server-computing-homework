package services

import (
	"net/http"
	"github.com/render"
	"html/template"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {
	homeTempl := template.Must(template.New("index.html").ParseFiles("templates/index.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		homeTempl.Execute(w, struct{}{})
	}
}