package main

import "fmt"

func search(arr []int, x int) int {
	for i := range arr {
		if arr[i] == x {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3}, 3))
}
