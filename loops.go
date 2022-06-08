package main

import "fmt"

func main() {

	// like while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// like traditional for loop
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
