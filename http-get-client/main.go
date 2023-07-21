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
type Occurrence struct {
	Words map[string]int `json:"words"`
}

func main() {
	args := os.Args
	myUrl := args[1]
	if len(args) < 2 {
		fmt.Printf("usage: ./http-get-client <url>\n")
		os.Exit(2)
	}
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(2)
	}
	response, err := http.Get(myUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
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
			log.Fatal(err)
		}
		fmt.Printf("JSON Parsed\nPage: %s\nWords: %s\n", words.Page, strings.Join(words.Words, ","))

	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			log.Fatal(err)
		}
		for word, occurrence := range occurrence.Words {
			fmt.Printf("%s: %d\n", word, occurrence)
		}
	default:
		fmt.Printf("Page not found \n")
	}

}
