package main

import "fmt"

type HttpRequest struct {
	URL     string
	Method  string
	Body    string
	Headers map[string]string
}

type HttpBuilder struct {
	req *HttpRequest
}

func NewHttpRequestBuilder() *HttpBuilder {
	return &HttpBuilder{req: &HttpRequest{Headers: make(map[string]string)}}
}

func (b *HttpBuilder) URL(url string) *HttpBuilder {
	b.req.URL = url
	return b
}

func (b *HttpBuilder) Method(method string) *HttpBuilder {
	b.req.Method = method
	return b
}

func (b *HttpBuilder) Body(body string) *HttpBuilder {
	b.req.Body = body
	return b
}

func (b *HttpBuilder) Header(key, value string) *HttpBuilder {
	b.req.Headers[key] = value
	return b
}

func (b *HttpBuilder) Build() (*HttpRequest, error) {
	if b.req.URL == "" {
		return nil, fmt.Errorf("Url should not be empty")
	}
	return b.req, nil
}

func main() {

	req, _ := NewHttpRequestBuilder().URL("http://localhost:9090").Method("GET").Body("").Header("K", "V").Build()
	fmt.Println(req.URL)
	fmt.Println(req.Method)
	fmt.Println(req.Headers)
}
