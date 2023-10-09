package client

import (
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"webscraper/internal/errors"
)

var defaultClient = &http.Client{}

func Get(url string) (string, error) {
	return getWithClient(url, defaultClient)
}

func getWithClient(url string, client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.NewError(errors.ErrCreatingGetRequest, "error creating get request to"+url)
	}

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.NewError(errors.ErrInGetRequest, "couldn't perform GET request to "+url)
	}
	defer resp.Body.Close()
	utf8Body, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(utf8Body)
	if err != nil {
		return "", errors.NewError(errors.ErrReadingResponse, "unable to read the response body")
	}
	return string(bytes), nil
}
