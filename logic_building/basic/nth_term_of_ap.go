package main

import "fmt"

func nthTermOfAp(a1 , a2 , n int) int{
	nthTerm := a1
	d := a2 - a1

	for i := 1 ; i < n ; i++ {

		nthTerm += d

	}

	return nthTerm
}

func nthTermOfApFormula(a1, a2, n int) int  {
	return (a1 + (n - 1) * (a2 - a1))
}

func main() {
	fmt.Println(nthTermOfAp(2,3,4))	
	fmt.Println(nthTermOfApFormula(2,3,4))	
}
