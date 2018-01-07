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
