package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	length := len(arr)

	for i := 0; i < length/2; i++ {
		var temp int

		temp = arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = temp
	}

	fmt.Println(arr)
}
