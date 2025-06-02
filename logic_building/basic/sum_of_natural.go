package main

import "fmt"

func findSum(n int) int {
	var res int

	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

func findSumFormula(n int) int {
	return n * (n + 1) / 2
}

func main() {
	fmt.Println(findSum(5))
	fmt.Println(findSumFormula(5))
}
