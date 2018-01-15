package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//DeleteUser .
func DeleteUser(password string) (bool, string) {
	ok, name, session := GetCurrentUser()
	if !ok {
		return false, "No login user"
	}
	url := URL + "/v1/users/" + name + "?password=" + password
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
	return DeleteRes(res.Body, res.StatusCode)
}

//DeleteRes .
func DeleteRes(resBody io.ReadCloser, statusCode int) (bool, string) {
	if statusCode == 204 {
		RemoveFile()
		return true, "delete user successfully"
	}
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "Fail to read body."
	}
	tmp := SingleMessageResponse{}
	if err := json.Unmarshal(body, &tmp); err != nil {
		fmt.Fprintln(os.Stderr, "Can not resolve body.")
	}
	return false, tmp.Message
}
