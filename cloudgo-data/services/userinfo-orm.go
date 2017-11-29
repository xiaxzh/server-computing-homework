package services

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/user/hello/entities"
	"strconv"
	"fmt"
)

func postUserInfoOrmHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ErrorIndo string} {"Bad Input!"})
			return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
		u.DepartName = req.Form["departname"][0]
		entities.UserInfoServiceOrmSave(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoOrmHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)
			fmt.Println(i)
			u := entities.UserInfoServiceOrmFindByID(int(i))
			formatter.JSON(w, http.StatusOK, u)
			return
		}
		ulist := entities.UserInfoServiceOrmFindAll()
		formatter.JSON(w, http.StatusOK, ulist)
	}
}