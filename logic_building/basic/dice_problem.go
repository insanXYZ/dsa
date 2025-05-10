package main

import "fmt"

func oppositeFaceOfDice(i int) int{
	return 7 - i 
}
//
// The idea is based on the observation that the sum of two opposite sides of a cubical dice is equal to 7. So, just subtract the given n from 7 and print the answer.
//
func main() {
	fmt.Println(oppositeFaceOfDice(1))
}
