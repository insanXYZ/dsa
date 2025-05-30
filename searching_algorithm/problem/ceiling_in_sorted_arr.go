package main

import "fmt"

// return index
func ceil_search(arr []int, x int) int {
	if x <= arr[0] {
		return 0
	}

	for i := range len(arr) - 1 {
		if arr[i] == x {
			return i
		}

		if arr[i] < x && arr[i+1] >= x {
			return i + 1
		}
	}

	return -1
}

func ceil_binary_search(arr []int, x int) int {

	low := 0
	high := len(arr) - 1
	res := -1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			res = mid
			high = mid - 1
		} else if arr[mid] < x {
			low = mid + 1
		}
	}

	return res

}

func main() {
	arr := []int{1, 2, 8, 10, 10, 12, 19}
	fmt.Println(arr[ceil_search(arr, 3)])
	fmt.Println(arr[ceil_binary_search(arr, 3)])
}
