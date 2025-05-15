package main

import "fmt"

func isPower(x, y int) bool {
	if x == 1 {
		return y == 1
	}

	pow := 1

	for pow < y {
		pow *= x
	}

	return pow == y
}

func main() {
	fmt.Println(isPower(10, 1))
}
