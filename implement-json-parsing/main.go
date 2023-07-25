package main

type ParsedStruct struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float64   `json:"percentages"`
	Special      []*string      `json:"special"`
	ExtraSpecial []any `json:"extraSpecial"`
}


// use ParseJson() to parse the JSON from the server and handle CLI flags
func main() {

}

// function for parsing JSON from the server based on the struct
func parseJson(a ParsedStruct) {

}
