package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: ./hello-world <argument>\n")
		os.Exit(1)
	}
	fmt.Printf("Hello world\nos.Args: %v\nArguments: %v\n", args, args[1])
}
