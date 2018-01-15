package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// FindUser .
func FindUser(id string) (bool, string, SingleUserInfo) {
	ok, _, session := GetCurrentUser()
	if !ok {
		return false, "No login user", SingleUserInfo{}
	}
	tarURL := URL + "/v1/users/" + id
	req, err := http.NewRequest("GET", tarURL, strings.NewReader(""))

	if err != nil {
		return false, "error : some mistakes happend in creating request to server", SingleUserInfo{}
	}
	req.Header.Set("Cookie", "key="+session)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, "error : some mistakes happend in sending request to server", SingleUserInfo{}
	}
	defer resp.Body.Close()
	return FindRes(resp.Body, resp.StatusCode)
}

// FindRes .
func FindRes(respBody io.ReadCloser, statusCode int) (bool, string, SingleUserInfo) {
	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		return false, "error : Some mistakes happend in forming body", SingleUserInfo{}
	}
	temp := UserKeyResponse{}
	if err := json.Unmarshal(body, &temp); err != nil {
		return false, "error : Some mistakes happend in parsing body", SingleUserInfo{}
	}
	if statusCode == 200 {
		return true, temp.Message, SingleUserInfo{temp.ID, temp.UserName, temp.Email, temp.Phone}
	}
	return false, temp.Message, SingleUserInfo{}
}
