package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	GetResponseOutput  *http.Response
	PostResponseOutput *http.Response
}

func (m MockClient) Get(url string) (resp *http.Response, err error) {
	return m.GetResponseOutput, nil
}

func (m MockClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return m.PostResponseOutput, nil
}
func TestHandleGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Errorf("marshal error: %s", err)
	}

	apiInstance := API{
		Options: Options{},
		Client: MockClient{
			GetResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.HandleGetRequest("http://localhost:8080/words")
	if err != nil {
		t.Errorf("handleGetResponse error: %s", err)
	}
	// if response == nil {
	// 	t.Fatalf("response is empty")
	// }
	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("unexpected response: %s", response.GetResponse())
	}
}
