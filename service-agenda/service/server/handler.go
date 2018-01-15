package server

import (
	"net/http"
	"strings"

	"github.com/freakkid/service-agenda/service/entities"
	"github.com/unrolled/render"
)

// get user key by username and password, need key
func userLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		sessionID, status, responseJSON := entities.AgendaService.LoginAndGetSessionID(req.FormValue("username"), req.FormValue("password"))
		if sessionID != "" || status == http.StatusOK { // login successfully and set cookie
			http.SetCookie(w, &http.Cookie{Name: "key", Value: sessionID})
		}
		formatter.JSON(w, status, responseJSON)
	}
}

// get user key by username and password, need key
func userLogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                // parsing the parameters
		cookie, _ := req.Cookie("key") // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			lastIndex := strings.LastIndex(req.URL.Path, "/")
			if lastIndex != -1 {
				status, responseJSON := entities.AgendaService.LogoutAndDeleteSessionID(cookie.Value, req.URL.Path[lastIndex+1:])
				formatter.JSON(w, status, responseJSON)

			} else {
				formatter.JSON(w, http.StatusBadRequest, entities.SingleMessageResponse{Message: "empty id"})
			}
		}
	}
}

// create a new user by username, password, email, password, no need key
func createUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		status, responseJSON := entities.AgendaService.CreateUser(req.Body)
		formatter.JSON(w, status, responseJSON)
	}
}

// delete a user by password, need key
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                // parsing the parameters
		cookie, _ := req.Cookie("key") // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			lastIndex := strings.LastIndex(req.URL.Path, "/")
			if lastIndex != -1 {
				status, responseJSON := entities.AgendaService.DeleteUserByPassword(cookie.Value, req.URL.Path[lastIndex+1:], req.FormValue("password"))
				formatter.JSON(w, status, responseJSON)

			} else {
				formatter.JSON(w, http.StatusBadRequest, entities.SingleMessageResponse{Message: "empty id"})
			}
		}
	}
}

// list limit users or get a user by id, need key
func usersInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                // parsing the parameters
		cookie, _ := req.Cookie("key") // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			status, responseJSON := entities.AgendaService.ListUsersByLimit(cookie.Value, req.FormValue("limit"), req.FormValue("offset"))
			formatter.JSON(w, status, responseJSON)
		}
	}
}

func userInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                // parsing the parameters
		cookie, _ := req.Cookie("key") // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			lastIndex := strings.LastIndex(req.URL.Path, "/")
			if lastIndex != -1 {
				status, responseJSON := entities.AgendaService.GetUserInfoByID(cookie.Value, req.URL.Path[lastIndex+1:])
				formatter.JSON(w, status, responseJSON)
			} else {
				formatter.JSON(w, http.StatusBadRequest, entities.SingleMessageResponse{Message: "empty id"})
			}
		}
	}
}

func changeUserPasswordHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                // parsing the parameters
		cookie, _ := req.Cookie("key") // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			lastIndex := strings.LastIndex(req.URL.Path, "/")
			if lastIndex != -1 {
				status, responseJSON := entities.AgendaService.ChangeUserPassword(cookie.Value, req.URL.Path[lastIndex+1:],
					req.FormValue("password"), req.FormValue("newpassword"), req.FormValue("confirmation"))
				formatter.JSON(w, status, responseJSON)
			} else {
				formatter.JSON(w, http.StatusBadRequest, entities.SingleMessageResponse{Message: "empty id"})
			}
		}
	}
}
