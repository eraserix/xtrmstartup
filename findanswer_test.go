package main

import "testing"

func TestFindName(t *testing.T) {
	name := findTheAnswer("what is your name")
	if name != "go" {
		t.Errorf("name not equel, expected go, found %v", name)
	}
}

func TestFindLargest(t *testing.T) {
	largest := findTheAnswer("which of the following numbers is the largest: 464, 54, 3, 8")
	if largest != "464" {
		t.Errorf("largest number was %v, expected 464", largest)
	}
}

func TestSimpleAddition(t *testing.T) {
	sum := findTheAnswer("what is 12 plus 16")
	if sum != "28" {
		t.Errorf("sum of 12 16 was %v, expected 28", sum)
	}
}

func TestSimpleMultiplication(t *testing.T) {
	product := findTheAnswer("what is 15 multiplied by 14")
	if product != "210" {
		t.Errorf("product wrong")
	}
}
