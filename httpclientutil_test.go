package httpclientutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"name": "test"})
	}))
	defer testServer.Close()

	type Test struct {
		Name string `json:"name"`
	}

	var obj Test
	err := GetJSON(testServer.URL, nil, &obj)
	if err != nil {
		t.Fatalf("Failed to get json: %v", err)
	}

	if obj.Name != "test" {
		t.Errorf("Expected name 'test', got '%v'", obj.Name)
	}
}

func TestPostJSON(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"name": "test"})
	}))
	defer testServer.Close()

	type Test struct {
		Name string `json:"name"`
	}

	var obj Test
	err := PostJSON(testServer.URL, nil, &Test{Name: "test"}, &obj)
	if err != nil {
		t.Fatalf("Failed to post json: %v", err)
	}

	if obj.Name != "test" {
		t.Errorf("Expected name 'test', got '%v'", obj.Name)
	}
}

func TestPutJSON(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"name": "test"})
	}))
	defer testServer.Close()

	type Test struct {
		Name string `json:"name"`
	}

	var obj Test
	err := PutJSON(testServer.URL, nil, &Test{Name: "test"}, &obj)
	if err != nil {
		t.Fatalf("Failed to put json: %v", err)
	}

	if obj.Name != "test" {
		t.Errorf("Expected name 'test', got '%v'", obj.Name)
	}
}

func TestDeleteJSON(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"name": "test"})
	}))
	defer testServer.Close()

	type Test struct {
		Name string `json:"name"`
	}

	var obj Test
	err := DeleteJSON(testServer.URL, nil, &obj)
	if err != nil {
		t.Fatalf("Failed to delete json: %v", err)
	}

	if obj.Name != "test" {
		t.Errorf("Expected name 'test', got '%v'", obj.Name)
	}
}

func TestGet(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	}))
	defer testServer.Close()

	respBody, err := Get(testServer.URL, nil)
	if err != nil {
		t.Fatalf("Failed to get: %v", err)
	}

	if string(respBody) != "test" {
		t.Errorf("Expected response body 'test', got '%v'", string(respBody))
	}
}
