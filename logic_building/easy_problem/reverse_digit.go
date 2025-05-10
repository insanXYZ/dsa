package main

import (
	"fmt"
	"strconv"
)

func reverseDigits(n int) int {
	var rev int

	for n != 0 {
		rev = rev * 10 + n % 10
		n /= 10
	}

	return rev
}

func reverseDigitsString(n int) int {
	
	itoa := strconv.Itoa(n)

	ritoa := []rune(itoa)

	for i := 0 ; i < len(itoa) / 2 ; i++ {
		ritoa[i] , ritoa[len(itoa) - 1 - i] = ritoa[len(itoa) - 1 - i] , ritoa[i]
	}

	atoi , _ := strconv.Atoi(string(ritoa))

	return atoi

}

func main(){
	fmt.Println(reverseDigits(122))
	fmt.Println(reverseDigitsString(1223))
}
