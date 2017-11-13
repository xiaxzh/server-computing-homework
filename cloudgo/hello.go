package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"encoding/json"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// get the lexme in the group --- {user}
	vars := mux.Vars(r)
	user := vars["user"]
	// construct Json --- "Test" : "Hello {user}"
	j, _ := json.Marshal(struct{ Test string }{ "Hello " + user })
	// add Json to Response
	w.Write([]byte(j))
	w.Write([]byte("\n\r"))
}

func main() {
	// use mux to produce a Router --- r
	r := mux.NewRouter()
	// add HandleFunc when request to "/"
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("enter '/' diretory"))
	})
	// add HandleFunc when request to "/hello/{user}"	
	r.HandleFunc("/hello/{user}", sayHello)
	// start server under negroni framework
	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":9090")
}