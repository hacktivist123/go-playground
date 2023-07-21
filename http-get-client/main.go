package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return  strings.Join(w.Words, ",")
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return strings.Join(out, ",")
}

type Response interface {
	GetResponse() string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: ./http-get-client <url>\n")
		os.Exit(2)
	}
	res, err := handleRequest(args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("No response \n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", res.GetResponse())
}

func handleRequest(requestURL string) (Response, error) {
	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, fmt.Errorf("validation error: %s", err)
	}
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("http-get error: %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, body)
	}
	var page Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		log.Fatal(err)
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %s", err)
		}
		return words, nil

	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %s", err)
		}
		return occurrence, nil
	}

	return nil, nil
}
