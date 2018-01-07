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

func findTheAnswer(question string) string {
	if strings.HasPrefix(question, "which of the following numbers is the largest") {
		return findLargest(question)
	} else if strings.HasPrefix(question, "what is your name") {
		return "go"
	} else if strings.HasPrefix(question, "what is") {
		return doSimpleCalculation(question)
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
