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

func TestSimpleSubtraction(t *testing.T) {
	difference := findTheAnswer("what is 4 minus 19")
	assert.Equal(t, "-15", difference)
}

func TestFindPrimes(t *testing.T) {
	primes := findTheAnswer("which of the following numbers are primes: 819, 401, 369, 383")
	assert.Equal(t, "401, 383", primes)
}

func TestFindFibonacci(t *testing.T) {
	fibo := findTheAnswer("what is the 21st number in the Fibonacci sequence")
	assert.Equal(t, "10946", fibo)
}

func TestFindSquarAndCube(t *testing.T) {
	squareCube := findTheAnswer("which of the following numbers is both a" +
		" square and a cube: 531441, 1225, 199, 479")
	assert.Equal(t, "531441", squareCube)
}
