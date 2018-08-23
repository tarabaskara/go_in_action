package main

import (
	"fmt"
)

func main() {
	// Declaring an array set to its zero value
	// var array [5]int

	// Declaring an array using array literal
	//array := [5]int{10, 20, 30, 40, 50}

	// Declaring an array with Go calculating size
	//array := [...]int{10, 20, 30, 40, 50}

	// Accessing array elements
	//array := [5]int{10, 20, 30, 40, 50}
	//array[1] = 35

	// Accessing array pointer of elements
	// array := [5]*int{0: new(int), 1: new(int)}
	// *array[0] = 10
	// *array[1] = 20

	// Assigning one array to another of the same type
	// var array1 [5]int
	// array2 := [5]int{1, 2, 3, 4, 5}
	// array1 = array2

	// Assigning one array of pointers to another
	// var array1 [3]*int
	// array2 := [3]*int{new(int), new(int), new(int)}
	// *array2[0] = 10
	// *array2[1] = 20
	// *array2[2] = 30
	// array1 = array2

	// Declaring two-dimensional array
	//var array [4][2]int

	// literal
	//array := [4][2]int{1: {10, 10}, 2: {20, 20}}

	// Accessing elements of two-dimensional array
	// var array [2][2]int
	// array[0][0] = 10
	// array[0][1] = 20
	// array[1][0] = 30
	// array[1][1] = 40

	// Assigning multidimensional arrays of the same type
	var array1 [2][2]int
	var array2 [2][2]int
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	array1 = array2

	fmt.Println(array1)

}
