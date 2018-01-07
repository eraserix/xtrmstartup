package main

import (
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findTheAnswer(question string) string {
	if strings.HasPrefix(question, "which of the following numbers is the largest") {
		return findLargest(question)
	} else if strings.HasPrefix(question, "what is your name") {
		return "go"
	}
	return ""
}

func findLargest(question string) string {
	numbers := strings.Split(strings.Split(question, ":")[1], ",")

	maxV := 0
	for _, n := range numbers {
		parsed, _ := strconv.Atoi(strings.TrimSpace(n))
		maxV = max(parsed, maxV)
	}
	return strconv.Itoa(maxV)
}
