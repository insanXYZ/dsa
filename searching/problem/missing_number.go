package main

import "fmt"

func missingNum(arr []int) int {
	n := len(arr) + 1

	for i := 1; i < n+1; i++ {

		found := false

		for j := range n - 1 {
			if arr[j] == i {
				found = true
				break
			}
		}

		if !found {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(missingNum([]int{20, 21, 22, 23, 25}))
}
