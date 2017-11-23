package services

import (
	"net/http"
	"os"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/render"
)

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.HandleFunc("/js", jsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/form", formHandler()).Methods("POST")
	mx.HandleFunc("/unknown", unknownHandler()).Methods("GET")
}

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options {
		Directory: "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}