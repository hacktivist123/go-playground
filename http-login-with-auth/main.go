package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	return strings.Join(w.Words, ",")
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
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "use a password to access the API")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("validation error: %s\n Usage: ./http-get-client -h", err)
		flag.Usage()
		os.Exit(1)
	}

	if password != "" {
		token, err := handleLoginRequest(parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			if requestErr, ok := err.(RequestError); ok {
				fmt.Printf("Error: %s (HTTP Code: %d, Body: %s\n)", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			}
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("token: %s\n", token)
		os.Exit(0)
	}

	res, err := handleRequest(parsedURL.String())
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
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
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
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurences unmarshal error: %s", err),
			}
		}
		return occurrence, nil
	}

	return nil, nil
}
