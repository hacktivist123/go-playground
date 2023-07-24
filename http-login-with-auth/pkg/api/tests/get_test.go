package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hacktivist123/go-playground/http-login-with-auth/pkg/api"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput, nil

}

func TestHandleGetRequest(t *testing.T) {
	words := api.WordsPage{
		Page: api.Page{Name: "words"},
		Words: api.Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}


	apiInstance := api.API{
		Options: api.Options{},
		Client: MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.HandleGetRequest("http://localhost/words")
	if err != nil {
		t.Errorf("handleGetResponse error: %s", err)
	}
	if response == nil {
		t.Fatalf("response is empty")
	}
	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("unexpected response: %s", response.GetResponse())
	}
}
