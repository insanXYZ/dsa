package main

import (
	"fmt"
	"slices"
)

type ChainHash struct {
	bucket int
	arrs   [][]int
}

func NewChainHash(v int) *ChainHash {
	return &ChainHash{
		bucket: v,
		arrs:   make([][]int, v),
	}
}

func (c *ChainHash) hashFunc(x int) int {
	return x % c.bucket
}

func (c *ChainHash) insertItem(v int) {
	index := c.hashFunc(v)
	c.arrs[index] = append(c.arrs[index], v)
}

func (c *ChainHash) deleteItem(v int) {
	index := c.hashFunc(v)

	if i := slices.Index(c.arrs[index], v); i != -1 {
		c.arrs[index] = slices.Concat(c.arrs[index][:i], c.arrs[index][i+1:])
	}
}

func main() {
	ch := NewChainHash(7)
	arrNum := []int{15, 11, 27, 8, 12}

	for _, v := range arrNum {
		ch.insertItem(v)
	}

	fmt.Println(ch.arrs)
}
