package tools

import "bytes"
import "net/http"

func POST(url string, contentType string, body *bytes.Buffer) (*http.Response, error) {
	resp, err := http.Post(url, contentType, body)
	return resp, err
}

func GET() {}
