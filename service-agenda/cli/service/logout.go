package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//Logout .
func Logout() (bool, string) {
	ok, name, session := GetCurrentUser()
	if !ok {
		return false, "No login user"
	}
	url := URL + "/v1/user/logout/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Cookie", "key="+session)
	if err != nil {
		return false, "Can not construct DELETE request."
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, "Send delete request failed."
	}
	defer res.Body.Close()
	return LogoutRes(res.Body, res.StatusCode)
}

//LogoutRes .
func LogoutRes(resBody io.ReadCloser, statusCode int) (bool, string) {
	if statusCode == 204 {
		RemoveFile()
		return true, "log out successfully"
	}

	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "Fail to read body."
	}
	tmp := SingleMessageResponse{}
	if err := json.Unmarshal(body, &tmp); err != nil {
		return false, "Can not resolve body."
	}
	return false, tmp.Message
}
