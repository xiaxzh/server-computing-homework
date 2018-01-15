package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer -- start server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	negroniInstance := negroni.Classic()
	muxInstance := mux.NewRouter()

	initRoutes(muxInstance, formatter)
	negroniInstance.UseHandler(muxInstance)

	return negroniInstance
}

func initRoutes(muxInstance *mux.Router, formatter *render.Render) {

	muxInstance.HandleFunc("/v1/user/login", userLoginHandler(formatter)).Methods("GET")
	muxInstance.PathPrefix("/v1/users/").Handler(userInfoHandler(formatter)).Methods("GET")
	muxInstance.PathPrefix("/v1/users/").Handler(deleteUserHandler(formatter)).Methods("DELETE")
	muxInstance.PathPrefix("/v1/users/").Handler(changeUserPasswordHandler(formatter)).Methods("PATCH")
	muxInstance.PathPrefix("/v1/user/logout").Handler(userLogoutHandler(formatter)).Methods("DELETE")
	muxInstance.HandleFunc("/v1/users", usersInfoHandler(formatter)).Methods("GET")
	muxInstance.HandleFunc("/v1/users", createUserHandler(formatter)).Methods("POST")

	muxInstance.NotFoundHandler = notImplementedHandler() // redirect 404 to 501
}

// handle /unknown  -- 501
func notImplementedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("501 - the request method is not supported by the server and cannot be handled!"))
	}
}
