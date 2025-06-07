package main

import (
	"fmt"
	"slices"
	"sort"
)

func sortMatrix(arr [][]int) {
	ints := slices.Concat(arr...)
	sort.Ints(ints)
	n := len(arr[0])

	for i := 0; i < len(arr); i++ {
		a := n * i
		b := a + n
		arr[i] = ints[a:b]
	}
}

func main() {
	arr := [][]int{
		[]int{5, 4, 7},
		[]int{1, 3, 8},
		[]int{2, 9, 6},
	}

	sortMatrix(arr)
	fmt.Println(arr)
}
