package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestFindRes(t *testing.T) {
	b, err := json.Marshal(&UserKeyResponse{Message: "get user info successfully", ID: 1, UserName: "hnx", Email: "email@qq.com", Phone: "12345678901"})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&UserKeyResponse{Message: "invalid id", ID: -1, UserName: "", Email: "", Phone: ""})
	checkErr(err1)
	p1 := ioutil.NopCloser(bytes.NewReader(b1))

	b2, err2 := json.Marshal(&UserKeyResponse{Message: "please re-login", ID: -1, UserName: "", Email: "", Phone: ""})
	checkErr(err2)
	p2 := ioutil.NopCloser(bytes.NewReader(b2))

	type args struct {
		respBody   io.ReadCloser
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
		want2 SingleUserInfo
	}{

		{
			name:  "GetUserInfoSucceed",
			args:  args{p, 200},
			want:  true,
			want1: "get user info successfully",
			want2: SingleUserInfo{ID: 1, UserName: "hnx", Email: "email@qq.com", Phone: "12345678901"},
		},
		{
			name:  "InvalidID",
			args:  args{p1, 400},
			want:  false,
			want1: "invalid id",
			want2: SingleUserInfo{},
		},
		{
			name:  "StatusUnauthorized",
			args:  args{p2, 401},
			want:  false,
			want1: "please re-login",
			want2: SingleUserInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := FindRes(tt.args.respBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("FindRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindRes() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("FindRes() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
