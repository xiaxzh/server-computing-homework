package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"encoding/json"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	j, _ := json.Marshal(struct{ Test string }{ "Hello " + user })
	w.Write([]byte(j))
	w.Write([]byte("\n\r"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("enter '/' diretory"))
	})
	r.HandleFunc("/hello/{user}", sayHello)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":9090")
}