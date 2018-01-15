package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func TestLogoutRes(t *testing.T) {
	b, err := json.Marshal(&SingleMessageResponse{Message: "log out successfully"})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&SingleMessageResponse{Message: "invalid id"})
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
	}{
		{
			name:  "logout successfully",
			args:  args{resBody: p, statusCode: 204},
			want:  true,
			want1: "log out successfully",
		},
		{
			name:  "InvalidID",
			args:  args{resBody: p1, statusCode: 400},
			want:  false,
			want1: "invalid id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LogoutRes(tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("LogoutRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LogoutRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
