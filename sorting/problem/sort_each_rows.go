package main

import (
	"fmt"
	"sort"
)

func sortRows(arr [][]int) {
	for _, v := range arr {
		sort.Ints(v)
	}
}

func main() {
	arrs := [][]int{
		[]int{77, 11, 22, 3},
		[]int{11, 89, 1, 12},
		[]int{32, 11, 56, 7},
		[]int{11, 22, 44, 33},
	}

	sortRows(arrs)
	fmt.Println(arrs)
}
