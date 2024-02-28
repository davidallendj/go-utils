package httpx

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type Header map[string]string
type Body []byte
type Method = string

const (
	METHOD_GET     Method = "GET"
	METHOD_POST    Method = "POST"
	METHOD_PUT     Method = "PUT"
	METHOD_PATCH   Method = "PATCH"
	METHOD_DELETE  Method = "DELETE"
	METHOD_HEAD    Method = "HEAD"
	METHOD_CONNECT Method = "CONNECT"
	METHOD_OPTIONS Method = "OPTIONS"
	METHOD_TRACE   Method = "TRACE"
)

func MakeHTTPRequest(url string, httpMethod Method, body Body, headers Header) (*http.Response, []byte, error) {
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
