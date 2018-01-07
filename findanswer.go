package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toInt(s []string) []int {
	ints := make([]int, len(s))
	for i, n := range s {
		ints[i], _ = strconv.Atoi(strings.TrimSpace(n))
	}
	return ints
}

func findTheAnswer(question string) string {
	switch {
	case strings.HasPrefix(question, "which of the following numbers is the largest"):
		return findLargest(question)
	case strings.HasPrefix(question, "what is your name"):
		return "go"
	case strings.HasPrefix(question, "what is"):
		return doSimpleCalculation(question)
	}
	return ""
}

func findLargest(question string) string {
	numbers := strings.Split(strings.Split(question, ":")[1], ",")
	ints := toInt(numbers)

	maxV := 0
	for _, n := range ints {
		maxV = max(n, maxV)
	}
	return strconv.Itoa(maxV)
}

func doSimpleCalculation(question string) string {
	var a, b int
	switch {
	case strings.Contains(question, "plus"):
		fmt.Sscanf(question, "what is %d plus %d", &a, &b)
		return strconv.Itoa(a + b)
	case strings.Contains(question, "multiplied by"):
		fmt.Sscanf(question, "what is %d multiplied by %d", &a, &b)
		return strconv.Itoa(a * b)
	}
	return ""
}
