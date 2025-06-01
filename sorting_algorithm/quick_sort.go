package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[high], arr[i] = arr[i], arr[high]

	return i
}

func anu(i int) {
	i++
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
