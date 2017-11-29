package services

import (
	"net/http"
	"html/template"
)

func formHandler() http.HandlerFunc {
	formTempl := template.Must(template.New("form.html").ParseFiles("templates/form.html"))
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		formTempl.Execute(w, struct{
			Header		map[string][]string
		} {Header: req.Header})
	}
}