package main

import (
	"fmt"
	"sort"
)

func getSecondLargest(arr []int) int {
	sort.Ints(arr)

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] != arr[i+1] {
			return arr[i]
		}
	}

	return -1
}

func main() {
	// arr := []int{12, 35, 1, 10, 34, 1}
	arr := []int{12, 35, 1, 10, 34, 1}
	fmt.Println(getSecondLargest(arr))
}
