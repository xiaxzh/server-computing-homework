package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CreateUser .
func CreateUser(createUsername string, createPassword string, createPhone string, createEmail string) (bool, string) {
	// regist user via http json
	reqBody := fmt.Sprintf("{\"username\": \"%v\", \"password\": \"%v\", \"phone\": \"%v\", \"email\": \"%v\"}", createUsername, createPassword, createPhone, createEmail)
	resp, err := http.Post(URL+"/v1/users", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()
	return CreateRes(resp.Body, resp.StatusCode)
}

// CreateRes .
func CreateRes(resBody io.ReadCloser, statusCode int) (bool, string) {
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		return false, "error : Some mistakes happend in reading resp.Body"
	}

	temp := CreateUserResponse{}
	if err = json.Unmarshal(body, &temp); err != nil {
		return false, "error : Some mistakes happend in parsing resp.Body"
	}
	return statusCode == http.StatusCreated, temp.Message
}
