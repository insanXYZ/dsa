package main

import (
	"fmt"
	"math"
)

func closestNumber(n , m int) int {
	q := n / m
	n1 :=  m * q

	var n2 int

	if n * m > 0 {
		n2 = m * (q + 1)
	} else {
		n2 = m * (q - 1)
	}

	if math.Abs(float64(n - n1)) < math.Abs(float64(n - n2)) {
		return n1
	}

	return n2
}

func main() {
	fmt.Println(closestNumber(13,4))
}
