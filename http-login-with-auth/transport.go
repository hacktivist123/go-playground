package main

import "net/http"

type JWTTransport struct {
	token     string
	transport http.RoundTripper
}

func (m JWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.token != "" {
		req.Header.Add("Authorization", "Bearer"+m.token)
	}
	return m.transport.RoundTrip(req)
}
