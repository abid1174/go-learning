package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	sort.Ints(arr)

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (right + left) / 2

		if arr[mid] == target {
			return arr[mid]
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 5, 7, 4, 9, 2, 6}
	var target int = 7

	fmt.Println(binarySearch(arr, target))
}
