package main

import "fmt"

func isEven(i int) bool {
	return i % 2 == 0
}

func isEvenBitwise(i int) bool {
	return i & 1 == 0
}

func isEvenShifting(i int) bool {
	return i == (i >> 1) << 1
}


func main() {
	
	a , b := 1 , 2

	fmt.Println(isEven(a) , isEven(b))
	fmt.Println(isEvenBitwise(a) , isEvenBitwise(b))
	fmt.Println(isEvenShifting(a) , isEvenShifting(b))
	
}
