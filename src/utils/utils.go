package utils

import (
	"io"
	"net/http"
)

func GetRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}
