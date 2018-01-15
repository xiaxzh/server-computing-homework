package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestListRes(t *testing.T) {
	b, err := json.Marshal(&UsersInfoResponse{Message: "invalid limit", SingleUserInfoList: []SingleUserInfo{}})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&UsersInfoResponse{Message: "please re-login", SingleUserInfoList: []SingleUserInfo{}})
	checkErr(err1)
	p1 := ioutil.NopCloser(bytes.NewReader(b1))

	type args struct {
		resBody    io.ReadCloser
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
		want2 []SingleUserInfo
	}{
		{
			name:  "InvalidLimit",
			args:  args{resBody: p, statusCode: 400},
			want:  false,
			want1: "invalid limit",
			want2: []SingleUserInfo{},
		},
		{
			name:  "StatusUnauthorized",
			args:  args{resBody: p1, statusCode: 401},
			want:  false,
			want1: "please re-login",
			want2: []SingleUserInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ListRes(tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("ListRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ListRes() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ListRes() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
