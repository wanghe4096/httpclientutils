package httpclientutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get sends an HTTP GET request with custom headers
func Get(url string, headers map[string]string) ([]byte, error) {
	return sendRequest(http.MethodGet, url, headers, nil)
}

// Post sends an HTTP POST request with custom headers and body
func Post(url string, headers map[string]string, body []byte) ([]byte, error) {
	return sendRequest(http.MethodPost, url, headers, body)
}

// Put sends an HTTP PUT request with custom headers and body
func Put(url string, headers map[string]string, body []byte) ([]byte, error) {
	return sendRequest(http.MethodPut, url, headers, body)
}

// Delete sends an HTTP DELETE request with custom headers
func Delete(url string, headers map[string]string) ([]byte, error) {
	return sendRequest(http.MethodDelete, url, headers, nil)
}

// sendRequest sends an HTTP request with custom headers and body
func sendRequest(method string, url string, headers map[string]string, body []byte) ([]byte, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a request object and set headers
	req, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %v", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response: %v", err)
	}

	return respBody, nil
}
