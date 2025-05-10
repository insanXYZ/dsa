package main

import "fmt"

func sumOfDigits(n int) int {
	var sum int

	for n != 0 {

		sum += n % 10

		n /= 10

	}

	return sum
}

func sumOfDigitsRecursive(n int) int {

	if n == 0 {
		return 0
	}

	return (n % 10) + sumOfDigitsRecursive(n/10)

}

func sumOfDigitsString(s string) int {

	var sum int

	for i := len(s) -1 ; i>= 0 ; i-- {

		a := s[i] - '0'

		sum += int(a)

	}

	return sum

}

func main(){
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigitsRecursive(12))
	fmt.Println(sumOfDigitsString("12"))
}
