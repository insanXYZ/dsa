package main

import "fmt"

func firstRepeatingElement(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := 1 + i; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return arr[i]
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(firstRepeatingElement([]int{10, 5, 3, 4, 3, 5, 6}))
}
