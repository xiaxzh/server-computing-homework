package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// ListAllUsers .
func ListAllUsers(limit string, offset string) (bool, string, []SingleUserInfo) {
	// list all user via http request
	ok, _, session := GetCurrentUser()
	if !ok {
		return false, "No login user", []SingleUserInfo{}
	}
	req, err := http.NewRequest("GET", URL+"/v1/users?limit="+limit+"&offset="+offset, strings.NewReader(""))
	if err != nil {
		return false, "error : some mistakes happend in creating request to server", []SingleUserInfo{}
	}
	req.Header.Set("Cookie", "key="+session)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, "error : some mistakes happend in sending request to server", []SingleUserInfo{}
	}
	defer resp.Body.Close()
	return ListRes(resp.Body, resp.StatusCode)
}

// ListRes .
func ListRes(resBody io.ReadCloser, statusCode int) (bool, string, []SingleUserInfo) {
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "error : some mistakes happend in forming body", []SingleUserInfo{}
	}
	temp := UsersInfoResponse{}
	if err := json.Unmarshal(body, &temp); err != nil {
		return false, "error : some mistakes happend in parsing body", []SingleUserInfo{}
	}

	if statusCode == http.StatusOK {
		ret := make([]SingleUserInfo, len(temp.SingleUserInfoList))
		for index, each := range temp.SingleUserInfoList {
			ret[index].ID = each.ID
			ret[index].UserName = each.UserName
			ret[index].Phone = each.Phone
			ret[index].Email = each.Email
		}
		return true, temp.Message, ret
	}
	return false, temp.Message, []SingleUserInfo{}
}
