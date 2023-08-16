package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFmap(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, mapSlice(func(x int) int { return x + 1 }, []int{0, 1, 2}))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 10, sum([]int{1, 2, 3, 4}))
}

func TestSecsSum(t *testing.T) {
	assert.Equal(t, 360, secsSum1([]int{1, 2, 3}))
	assert.Equal(t, 360, secsSum2([]int{1, 2, 3}))
}

func TestNilTreeLookup(t *testing.T) {
	var tree *BinaryTree[string, int] = nil
	assert.Nil(t, tree.Lookup("a"))
}
