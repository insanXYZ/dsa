package main

import "fmt"

func largest(arr []int) int {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func main() {
	arr := []int{10, 324, 45, 90, 9808}
	fmt.Println(largest(arr))
}
