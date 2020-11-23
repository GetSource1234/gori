package client

import (
	"io/ioutil"
	"net/http"
)

// Request wrapper to initiate HTTP request
func Request(URL string) (
	status int, body []byte, resHeader http.Header, err error,
) {

	// Create an HTTP Request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return status, body, resHeader, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}

	return

}
