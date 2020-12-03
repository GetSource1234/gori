package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	u "net/url"
	"strings"
	"sync"
	"time"
)

// Call HTTP request
func Call(
	uri string,
	wg *sync.WaitGroup,
	schemes []string,
	paths map[string][]string,
	tor bool,
	mu *sync.Mutex,
	errors *[]string,
) {

	for _, s := range schemes {
		for p, r := range paths {
			URL := s + uri + p
			status, body, _, err := request(URL, tor)
			if status == 200 {
				for _, v := range r {
					if strings.Contains(string(body), v) {
						fmt.Println("Leaked!!! - ", URL)
					}
				}
			}
			if err != nil {
				mu.Lock()
				*errors = append(*errors, fmt.Sprintf("error within : %s", err))
				mu.Unlock()
			}
			wg.Done()
		}
	}
}

func request(URL string, tor bool) (
	status int, body []byte, resHeader http.Header, err error,
) {

	// Create an HTTP Request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return status, body, resHeader, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	if tor {
		proxyURL := &u.URL{Scheme: "socks5", Host: "127.0.0.1:9150"}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}

	return resp.StatusCode, body, resp.Header, err
}
