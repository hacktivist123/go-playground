package main

import "fmt"

func main() {
	var arr1 [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("total: %v\n", arr1[1]+arr1[6])
	fmt.Printf("Length of the array: %d\nCapacity of the array: %d\n", len(arr1), cap(arr1))
	var arr2 []int = arr1[1:3]
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr2), cap(arr2))

}
