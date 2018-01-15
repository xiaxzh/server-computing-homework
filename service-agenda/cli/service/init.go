package service

import (
	"os"
)

//UserFile .
var UserFile string

//SessionFile .
var SessionFile string

// UserMap .
var UserMap string

// URL .
var URL string

// LoginRetJSON .
type LoginRetJSON struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// FindUserRetJSON .
type FindUserRetJSON struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// User .
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// UserKeyResponse -- GetUserByKeyAndID
type UserKeyResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// SingleMessageResponse -- DeleteUserByKeyAndPassword/ChangeUserPassword/GetUserKey
type SingleMessageResponse struct {
	Message string `json:"message"`
}

// SingleUserInfo .
type SingleUserInfo struct {
	ID       int
	UserName string
	Email    string
	Phone    string
}

// UsersInfoResponse -- ListUsersByKeyAndLimit
type UsersInfoResponse struct {
	Message            string
	SingleUserInfoList []SingleUserInfo
}

// CreateUserResponse -- CreateUser
type CreateUserResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

func init() {
	UserFile = "./currentUser"
	SessionFile = "./session"
	// UserMap = "./userMap"
	envURL := os.Getenv("SERVER_ADDR")
	PORT := os.Getenv("PORT")
	if len(envURL) == 0 {
		URL = "http://localhost:8080"
	} else {
		URL = "http://" + envURL + ":" + PORT
	}
}
