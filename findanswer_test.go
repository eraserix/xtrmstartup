package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheAnswer(t *testing.T) {
	cases := []struct {
		question, answer string
	}{
		{"what is your name", "go"},
		{"which of the following numbers is the largest: 464, 54, 3, 8", "464"},
		{"what is 12 plus 16", "28"},
		{"what is 15 multiplied by 14", "210"},
		{"what is 4 minus 19", "-15"},
		{"which of the following numbers are primes: 819, 401, 369, 383", "401, 383"},
		{"what is the 21st number in the Fibonacci sequence", "10946"},
		{"which of the following numbers is both a square and a cube: 531441, 1225, 199, 479",
			"531441"},
		{"what is 19 to the power of 18", "104127350297911241532841"},
		{"what is 19 to the power of 0", "1"},
		{"what is 14 plus 4 multiplied by 2", "22"},
		{"what is 16 plus 12 plus 2", "30"},
		{"what is the english scrabble score of ruby", "9"},
		{"which of the following is an anagram of \"dictionary\": " +
			"abdication, indicatory, butterfly, incendiary", "indicatory"},
	}

	for _, c := range cases {
		answer := findTheAnswer(c.question)
		assert.Equal(t, c.answer, answer)
	}
}
