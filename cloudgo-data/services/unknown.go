package services

import (
	"net/http"
)

func unknownHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}
}