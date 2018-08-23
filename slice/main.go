package main

import (
	"fmt"
)

func main() {
	// Declaring a slice integers by length and capacity
	//slice := make([]string, 3, 5)

	// Declaring a slice with slice literal
	//slice := []int{10, 20, 30}

	// Declaring a nil slice
	//var slice []int

	// Taking a slice of a slice
	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := slice[1:3]

	// Potential consequence of making changes to a slice
	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := slice[1:3]
	// newSlice[1] = 35

	// Using append to add an element to a slice
	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := slice[1:3]
	// newSlice = append(newSlice, 60)

	// Performing a three-index slice
	// source := []string{"a", "b", "c", "d", "e"}
	// slice := source[2:3:4]
	// slice = append(slice, "k")

	// Benefits of setting length and capacity to be the same
	// source := []string{"a", "b", "c", "d", "e"}
	// slice := source[2:3:3]
	// slice = append(slice, "k")

	// Iterating over a slice using for range
	// slice := []int{1, 2, 3, 4}
	// for index, value := range slice {
	// 	fmt.Printf("Index: %d Value: %d\n", index, value)
	// }

	// Using the blank identifier to ignore the index value
	slice := []int{1, 2, 3, 4}
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}

	// fmt.Println(source)
	// fmt.Println(slice)
}
