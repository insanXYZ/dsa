package main

import "fmt"

func merge(x, y []int) []int {
	var res []int

	i, j := 0, 0

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			res = append(res, x[i])
			i++
		} else {
			res = append(res, y[j])
			j++
		}
	}

	for ; i < len(x); i++ {
		res = append(res, x[i])
	}

	for ; j < len(y); j++ {
		res = append(res, y[j])
	}

	return res
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	l := mergeSort(arr[:len(arr)/2])
	r := mergeSort(arr[len(arr)/2:])

	return merge(l, r)
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(mergeSort(arr))
}
