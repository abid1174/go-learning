package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 2
	// i := 2

	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")
	}

	fmt.Println("===============")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	fmt.Println("===============")
}
