package httpx

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type HttpHeaders map[string]string
type HttpBody []byte
type HttpMethod = string

const (
	HTTP_METHOD_GET     HttpMethod = "GET"
	HTTP_METHOD_POST    HttpMethod = "POST"
	HTTP_METHOD_PUT     HttpMethod = "PUT"
	HTTP_METHOD_PATCH   HttpMethod = "PATCH"
	HTTP_METHOD_DELETE  HttpMethod = "DELETE"
	HTTP_METHOD_HEAD    HttpMethod = "HEAD"
	HTTP_METHOD_CONNECT HttpMethod = "CONNECT"
	HTTP_METHOD_OPTIONS HttpMethod = "OPTIONS"
	HTTP_METHOD_TRACE   HttpMethod = "TRACE"
)

func MakeHTTPRequest(url string, httpMethod HttpMethod, body HttpBody, headers HttpHeaders) (*http.Response, []byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, fmt.Errorf("could not create new HTTP request: %v", err)
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("could not make request: %v", err)
	}
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("could not read response body: %v", err)
	}
	return res, b, err
}
