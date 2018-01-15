package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//UpdatePassword .
func UpdatePassword(old, new, confirm string) (bool, string) {
	ok, name, session := GetCurrentUser()
	if !ok {
		return false, "No login user"
	}
	url := URL + "/v1/users/" + name + "?password=" + old + "&newpassword=" + new + "&confirmation=" + confirm
	req, err := http.NewRequest("PATCH", url, nil)
	req.Header.Set("Cookie", "key="+session)
	if err != nil {
		return false, "Can not construct PATCH request."
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, "Send patch request failed."
	}
	defer res.Body.Close()
	return UpdateRes(res.Body, res.StatusCode)
}

//UpdateRes .
func UpdateRes(resBody io.ReadCloser, statusCode int) (bool, string) {
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "Fail to read body."
	}
	tmp := SingleMessageResponse{}
	if err := json.Unmarshal(body, &tmp); err != nil {
		return false, "Can not resolve body."
	}
	if statusCode == 200 {
		return true, tmp.Message
	}
	return false, tmp.Message
}
