# httpclientutils

## HTTP Client Util

HTTP Client Util is a Go library for sending HTTP requests with custom headers.

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

## License

HTTP Client Util is licensed under the [MIT License](https://chat.openai.com/chat/LICENSE).