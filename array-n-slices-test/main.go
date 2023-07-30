package main

import "fmt"

func main() {
	var arr1 [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr1)
	fmt.Printf("total: %v\n", arr1[1]+arr1[6])
	fmt.Printf("Length of the array: %d\nCapacity of the array: %d\n", len(arr1), cap(arr1))
	var arr2 []int = arr1[1:3]
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr2), cap(arr2))
	for k := range arr2 {
		arr2[k] += 1
	}
	fmt.Println(arr2)
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr2), cap(arr2))
	fmt.Println(arr1)

	var arr3 []int = []int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 4)
	fmt.Println(arr3)
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr3), cap(arr3))

	var arr4 = make([]int, 3, 9)
	fmt.Println(arr4)
	fmt.Printf("Length of the Slice: %d\nCapacity of the Slice: %d\n", len(arr4), cap(arr4))

}
