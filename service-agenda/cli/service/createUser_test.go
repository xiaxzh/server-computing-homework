package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCreateRes(t *testing.T) {
	// for test
	b, err := json.Marshal(&CreateUserResponse{Message: "create user successfully", ID: 1, UserName: "hnx", Email: "email@qq.com", Phone: "12345678901"})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&CreateUserResponse{Message: "duplicate username", ID: -1, UserName: "", Email: "", Phone: ""})
	checkErr(err1)
	p1 := ioutil.NopCloser(bytes.NewReader(b1))

	b2, err2 := json.Marshal(&CreateUserResponse{Message: "empty input", ID: -1, UserName: "", Email: "", Phone: ""})
	checkErr(err2)
	p2 := ioutil.NopCloser(bytes.NewReader(b2))

	type args struct {
		resBody    io.ReadCloser
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// for test
		{
			name:  "CreateUserSuceed",
			args:  args{resBody: p, statusCode: 201},
			want:  true,
			want1: "create user successfully",
		},
		{
			name:  "DuplicateUsername",
			args:  args{resBody: p1, statusCode: 400},
			want:  false,
			want1: "duplicate username",
		},
		{
			name:  "EmptyInput",
			args:  args{resBody: p2, statusCode: 400},
			want:  false,
			want1: "empty input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateRes(tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("CreateRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
