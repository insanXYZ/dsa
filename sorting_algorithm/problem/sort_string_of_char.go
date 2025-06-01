package main

import (
	"fmt"
)

// use bubble sort
func sortString(s []byte) []byte {

	for i := 0; i < len(s)-1; i++ {
		swapped := false

		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				swapped = true
				s[j+1], s[j] = s[j], s[j+1]
			}
		}

		if !swapped {
			break
		}
	}

	return s
}

func main() {
	s := "insan"
	res := sortString([]byte(s))
	fmt.Println(string(res))
}
