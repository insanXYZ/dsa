package main

import "fmt"

func main() {
	var temp int
	a , b := 1 , 2

	fmt.Println(a,  b)

	temp = a
	a = b
	b = temp

	fmt.Println(a,  b)
}
