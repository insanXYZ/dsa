package main

import "fmt"

func factorial(n int) int {
	res := 1

	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}

func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorialRecursive(5))
}
