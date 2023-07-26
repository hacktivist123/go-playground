package main

import "fmt"

func main() {
	a := "string"
	testPointer(&a)
	fmt.Printf("a: %s \n", a)
}
func testPointer(a *string) {
	*a = "another string"
}
