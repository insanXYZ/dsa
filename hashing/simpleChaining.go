package hashing

import (
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

func (c *ChainHash) InsertItem(v int) {
	index := c.hashFunc(v)
	c.arrs[index] = append(c.arrs[index], v)
}

func (c *ChainHash) DeleteItem(v int) {
	index := c.hashFunc(v)

	if i := slices.Index(c.arrs[index], v); i != -1 {
		c.arrs[index] = slices.Concat(c.arrs[index][:i], c.arrs[index][i+1:])
	}
}
