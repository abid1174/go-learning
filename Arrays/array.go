package main

import (
	"fmt"
)

func main() {
	// create an integer array with size 5
	// and not initializing the values
	var a [5]int
	fmt.Println("New Array := ", a)

	// insert value in any index
	a[4] = 100
	fmt.Println("After insert value :=", a)

	// size of the array
	fmt.Println("Array size := ", len(a))

	// create an array with initial values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b :=", b)

	// create a two dimentional array and insert values
	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println("2d: ", d)
}
