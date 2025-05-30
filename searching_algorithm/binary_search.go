package main

import "fmt"

func binarySearch(arr []int, low, high, x int) int {

	for low <= high {

		mid := low + (high-low)/2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return -1
}

func main() {
	intSorted := []int{2, 3, 4, 10, 40}

	fmt.Println(binarySearch(intSorted, 0, len(intSorted)-1, 10))

}
