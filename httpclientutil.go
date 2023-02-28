package httpclientutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)
// GetJSON sends an HTTP GET request and unmarshals the response body into a JSON object
func GetJSON(url string, headers map[string]string, obj interface{}) error {
	respBody, err := Get(url, headers)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBody, obj); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}
	return nil
}

// PostJSON sends an HTTP POST request and unmarshals the response body into a JSON object
func PostJSON(url string, headers map[string]string, body interface{}, obj interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	respBody, err := Post(url, headers, bodyBytes)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBody, obj); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}
	return nil
}

// PutJSON sends an HTTP PUT request and unmarshals the response body into a JSON object
func PutJSON(url string, headers map[string]string, body interface{}, obj interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	respBody, err := Put(url, headers, bodyBytes)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBody, obj); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}
	return nil
}

// DeleteJSON sends an HTTP DELETE request and unmarshals the response body into a JSON object
func DeleteJSON(url string, headers map[string]string, obj interface{}) error {
	respBody, err := Delete(url, headers)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBody, obj); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}
	return nil
}



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
