package main

import "fmt"

func arrSortedOrNot(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{20, 23, 23, 45, 78, 88}

	fmt.Println(arrSortedOrNot(arr))
}
