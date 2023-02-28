# httpclientutils

## HTTP Client Util

HTTP Client Util is a Go library for sending HTTP requests with custom headers.

This is a package that provides a set of functions for making HTTP requests with custom headers and request bodies, and unmarshaling the response bodies into JSON objects. It uses the net/http package for making HTTP requests and the encoding/json package for marshaling and unmarshaling JSON.
## Features

Supports sending GET, POST, PUT, and DELETE requests.
Allows users to set custom headers for the request.
Returns the response body as a byte array.
Provides a simple API for sending HTTP requests.

## Usage

To use HTTP Client Util in your Go program, import the httpclientutil package:

```go
import "github.com/yourusername/httpclientutil"
```

Then, use one of the functions provided by the package to send an HTTP request with custom headers:

```go
// Send an HTTP GET request with custom headers
headers := map[string]string{
	"User-Agent": "Mozilla/5.0",
	"Accept":     "application/json",
}
response, err := httpclientutil.Get("https://example.com", headers)

// Send an HTTP POST request with custom headers and body
headers = map[string]string{
	"Content-Type": "application/json",
}
body := []byte(`{"name": "John", "age": 30}`)
response, err = httpclientutil.Post("https://example.com/api/users", headers, body)

```


## GET Requests
To make a GET request and unmarshal the response body into a JSON object, use the GetJSON function:

```go
url := "https://jsonplaceholder.typicode.com/todos/1"
headers := map[string]string{"Content-Type": "application/json"}
var todo Todo
err := httpclient.GetJSON(url, headers, &todo)

```

## POST Requests
To make a POST request and unmarshal the response body into a JSON object, use the PostJSON function:

```go
url := "https://jsonplaceholder.typicode.com/todos"
headers := map[string]string{"Content-Type": "application/json"}
todo := Todo{Title: "Buy milk", Completed: false}
body, _ := json.Marshal(todo)
var newTodo Todo
err := httpclient.PostJSON(url, headers, body, &newTodo)

```

## PUT Requests
To make a PUT request and unmarshal the response body into a JSON object, use the PutJSON function:
```go
url := "https://jsonplaceholder.typicode.com/todos/1"
headers := map[string]string{"Content-Type": "application/json"}
todo := Todo{ID: 1, Title: "Buy milk", Completed: true}
body, _ := json.Marshal(todo)
var updatedTodo Todo
err := httpclient.PutJSON(url, headers, body, &updatedTodo)

```

## DELETE Requests
To make a DELETE request and unmarshal the response body into a JSON object, use the DeleteJSON function:

```go
url := "https://jsonplaceholder.typicode.com/todos/1"
headers := map[string]string{"Content-Type": "application/json"}
var todo Todo
err := httpclient.DeleteJSON(url, headers, &todo)

```

## Custom Requests
To make a custom HTTP request with a method other than GET, POST, PUT, or DELETE, use the sendRequest function:

```go
url := "https://jsonplaceholder.typicode.com/todos/1"
headers := map[string]string{"Content-Type": "application/json"}
method := "PATCH"
todo := Todo{ID: 1, Title: "Buy milk", Completed: true}
body, _ := json.Marshal(todo)
var updatedTodo Todo
respBody, err := httpclient.sendRequest(method, url, headers, body)

```

## Functions

### GetJSON(url string, headers map[string]string, obj interface{}) error
Sends an HTTP GET request and unmarshals the response body into a JSON object.

### PostJSON(url string, headers map[string]string, body []byte, obj interface{}) error
Sends an HTTP POST request and unmarshals the response body into a JSON object.

### PutJSON(url string, headers map[string]string, body []byte, obj interface{}) error
Sends an HTTP PUT request and unmarshals the response body into a JSON object.

### DeleteJSON(url string, headers map[string]string, obj interface{}) error
Sends an HTTP DELETE request and unmarshals the response body into a JSON object.

### Get(url string, headers map[string]string) ([]byte, error)
Sends an HTTP GET request with custom headers.

### Post sends an HTTP POST request with custom headers and body
func Post(url string, headers map[string]string, body []byte)

### Put sends an HTTP PUT request with custom headers and body
func Put(url string, headers map[string]string, body []byte)

### Delete sends an HTTP DELETE request with custom headers
func Delete(url string, headers map[string]string)


## License

HTTP Client Util is licensed under the MIT License