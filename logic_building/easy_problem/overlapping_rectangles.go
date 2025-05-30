package main

import "fmt"

type Point struct {
	x, y int
}

func doOverlap(l1, r1, l2, r2 Point) bool {
	return l1.x < r2.x || l2.x < r1.x || r1.y < l2.y || r2.y < l1.y
}

func main() {
	l1 := Point{0, 10}
	r1 := Point{10, 0}
	l2 := Point{5, 5}
	r2 := Point{15, 0}

	fmt.Println(doOverlap(l1, r1, l2, r2))
}
