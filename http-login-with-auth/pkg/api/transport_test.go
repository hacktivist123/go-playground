package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type MockRoundTripper struct {
	RoundTripperOutput *http.Response
}

func (m MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") != "Bearer abc" {
		return nil, fmt.Errorf("wrong authorization header: %s", req.Header.Get("Authorization"))
	}
	return m.RoundTripperOutput, nil
}

func TestRoundTrip(t *testing.T) {
	loginResponse := LoginResponse{
		Token: "abc",
	}

	loginResponseBytes, err := json.Marshal(loginResponse)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}

	JWTTransport := JWTTransport{
		transport: MockRoundTripper{
			RoundTripperOutput: &http.Response{
				StatusCode: 200,
			},
		},
		httpClient: MockClient{
			PostResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(loginResponseBytes)),
			},
		},
		password: "xyz",
	}
	req := &http.Request{
		Header: make(http.Header),
	}
	res, err := JWTTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip error: %s", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("StatusCode is not 200: %d", res.StatusCode)
	}
}
