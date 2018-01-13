package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindName(t *testing.T) {
	name := findTheAnswer("what is your name")
	assert.Equal(t, "go", name)
}

func TestFindLargest(t *testing.T) {
	largest := findTheAnswer("which of the following numbers is the largest: 464, 54, 3, 8")
	assert.Equal(t, "464", largest)
}

func TestSimpleAddition(t *testing.T) {
	sum := findTheAnswer("what is 12 plus 16")
	assert.Equal(t, "28", sum)
}

func TestSimpleMultiplication(t *testing.T) {
	product := findTheAnswer("what is 15 multiplied by 14")
	assert.Equal(t, "210", product)
}
