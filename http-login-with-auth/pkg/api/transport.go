package api

import "net/http"

type JWTTransport struct {
	token     string
	transport http.RoundTripper
	password  string
	loginURL  string
}

func (m JWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.token == "" {
		if m.password != "" {
			token, err := handleLoginRequest(http.Client{}, m.loginURL, m.password)
			if err != nil {
				if err != nil {
					return nil, err
				}
			}
			m.token = token
		}
	}
	if m.token != "" {
		req.Header.Add("Authorization", "Bearer "+m.token)
	}
	return m.transport.RoundTrip(req)
}
