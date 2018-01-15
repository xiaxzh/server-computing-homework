package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func TestGetUserKeyRes(t *testing.T) {
	b, err := json.Marshal(&LoginRetJSON{Message: "login successfully", ID: 1})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&LoginRetJSON{Message: "incorrect username or password", ID: -1})
	checkErr(err1)
	p1 := ioutil.NopCloser(bytes.NewReader(b1))
	type args struct {
		sessionID  string
		resBody    io.ReadCloser
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name:  "LoginSucceed",
			args:  args{sessionID: "2rdtfygbhunjimko234", resBody: p, statusCode: 200},
			want:  true,
			want1: "login successfully",
		},
		{
			name:  "IncorrectUsernameAndPassword",
			args:  args{sessionID: "", resBody: p1, statusCode: 401},
			want:  false,
			want1: "incorrect username or password",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUserKeyRes(tt.args.sessionID, tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("GetUserKeyRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetUserKeyRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
