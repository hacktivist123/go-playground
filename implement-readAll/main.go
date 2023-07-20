package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	Content string
	Pos     int
}

func (m MySlowReader) Read(p []byte) (n int, err error) {
	if m.Pos+1 <= len(m.Content) {
		n := copy(p, []byte(m.Content[m.Pos:m.Pos+1]))
		m.Pos++
		return n, nil
	}
 return 0, io.EOF
}

func main() {

	MySlowReaderInstance := &MySlowReader{
		Content: "hello world",
	}

	output, err := io.ReadAll(MySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output: %s\n", output)
}
