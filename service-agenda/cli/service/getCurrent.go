package service

import (
	"io/ioutil"
	"os"
)

// GetCurrentUser .
func GetCurrentUser() (bool, string, string) {
	UserItem, err := ioutil.ReadFile(UserFile)
	if err != nil {
		return false, "", ""
	}
	SessionItem, err := ioutil.ReadFile(SessionFile)
	if err != nil {
		return false, "", ""
	}
	username := string(UserItem)
	session := string(SessionItem)
	if err != nil || len(username) == 0 || len(session) == 0 {
		return false, "", ""
	}
	return true, username, session
}

// RemoveFile .
func RemoveFile() {
	os.Remove(UserMap)
	os.Remove(UserFile)
	os.Remove(SessionFile)
}
