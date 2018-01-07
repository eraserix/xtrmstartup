package main

import "testing"

func TestFindName(t *testing.T) {
	name := findTheAnswer("what is your name")
	if name != "go" {
		t.Errorf("name not equel, expected go, found %v", name)
	}
}
