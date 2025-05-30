package main

import "fmt"

// (a + b) > c
// (a + c) > b
// (b + c) > a
func checkValidity(a, b, c int) bool {
	return (a+b) > c || (a+c) > b || (b+c) > a
}

func main() {
	fmt.Println(checkValidity(7, 10, 5))
}
