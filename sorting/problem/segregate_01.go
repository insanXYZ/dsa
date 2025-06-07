package main

import (
	"fmt"
)

func segregate0and1(arr []int) []int {

	res := make([]int, len(arr))

	for i, j := 0, len(arr)-1; i < len(arr); i++ {
		if arr[i] == 1 {
			res[j] = 1
			j--
		}
	}

	return res

}

func main() {
	arr := []int{0, 1, 0, 1, 0, 0, 1, 1, 1, 0}
	res := segregate0and1(arr)
	fmt.Println(res)
}
