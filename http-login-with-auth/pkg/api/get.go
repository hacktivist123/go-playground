package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Response interface {
	GetResponse() string
}

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type WordsPage struct {
	Page
	Words
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("Words: %s", strings.Join(w.Words, ", "))
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	words := []string{}
	for word, occurrence := range o.Words {
		words = append(words, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("Words: %s", strings.Join(words, ", "))
}

func (a API) HandleGetRequest(requestURL string) (Response, error) {

	response, err := a.Client.Get(requestURL)

	if err != nil {
		return nil, fmt.Errorf("get error: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, string(body))
	}

	if !json.Valid(body) {
		return nil, RequestError{
			Err:      "response is not a json",
			HTTPCode: response.StatusCode,
			Body:     string(body),
		}
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestError{
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
			HTTPCode: response.StatusCode,
			Body:     string(body),
		}
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Words unmarshal error: %s", err),
			}
		}

		return words, nil

	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, fmt.Errorf("Occurrence unmarshal error: %s", err)
		}

		return occurrence, nil
	}

	return nil, nil
}
