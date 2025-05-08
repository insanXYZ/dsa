package main

import "fmt"

func printTable(n int) {
	for i := range 11 {
		fmt.Printf("%v * %v = %v\n",n,i,n*i )
	}
}

func main(){
	printTable(5)
}
