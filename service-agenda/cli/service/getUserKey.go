package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)


// GetUserKey .
func GetUserKey(username string, password string) (bool, string) {
	tarURL := URL + "/v1/user/login?username=" + username + "&password=" + password
	resp, err := http.Get(tarURL)
	if err != nil {
		return false, "error : Some mistakes happend in sending get request to tarUrl"
	}
	defer resp.Body.Close()
	sessionID := ""
	for _, item := range resp.Cookies() {
		if item.Name == "key" {
			sessionID = item.Value
		}
	}
	return GetUserKeyRes(sessionID, resp.Body, resp.StatusCode)
}

// GetUserKeyRes .
func GetUserKeyRes(sessionID string, resBody io.ReadCloser, statusCode int) (bool, string) {
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "error : Some mistakes happend in reading body"
	}
	temp := LoginRetJSON{}
	if err = json.Unmarshal(body, &temp); err != nil {
		return false, "error : Some mistakes happend in parsing body"
	}
	if statusCode == http.StatusOK {

		if sessionID == "" {
			return false, "error : sessionID should not be empty"
		}
		// write to file -- user
		err = ioutil.WriteFile(UserFile, []byte(strconv.Itoa(temp.ID)), 0655)
		if err != nil {
			return false, "Some mistakes happend in writing to current user"
		}
		// write to file -- sessionID
		err = ioutil.WriteFile(SessionFile, []byte(sessionID), 0655)
		if err != nil {
			return false, "Some mistakes happend in writing to sessionID"
		}
		return true, temp.Message
	}
	return false, temp.Message
}
