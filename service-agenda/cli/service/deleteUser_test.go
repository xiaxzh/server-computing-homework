package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func TestDeleteRes(t *testing.T) {
	b, err := json.Marshal(&SingleMessageResponse{Message: "delete user successfully"})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	b1, err1 := json.Marshal(&SingleMessageResponse{Message: "incorrect password"})
	checkErr(err1)
	p1 := ioutil.NopCloser(bytes.NewReader(b1))

	b2, err2 := json.Marshal(&SingleMessageResponse{Message: "invalid id"})
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
		{name: "success",
			args:  args{resBody: p, statusCode: 204},
			want:  true,
			want1: "delete user successfully",
		},
		{name: "IncorrectPassword",
			args:  args{resBody: p1, statusCode: 401},
			want:  false,
			want1: "incorrect password",
		},
		{name: "InvalidID",
			args:  args{resBody: p2, statusCode: 400},
			want:  false,
			want1: "invalid id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DeleteRes(tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("DeleteRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeleteRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
