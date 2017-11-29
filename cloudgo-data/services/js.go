package services

import (
	"net/http"
	"github.com/unrolled/render"
)

func jsHandler(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Message string
		} {Message: "Here is the sample js."})
	}
}