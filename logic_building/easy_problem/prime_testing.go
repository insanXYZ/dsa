package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%2 == 0 {
			return false
		}
	}

	return true

}

func isPrimeSqrt(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%2 == 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(isPrime(8))
	fmt.Println(isPrimeSqrt(8))
}
