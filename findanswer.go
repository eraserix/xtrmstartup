package main

import (
	"fmt"
	"math"
	"regexp"
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

func getIntList(question string) []int {
	numbers := strings.Split(strings.Split(question, ":")[1], ",")
	return toInt(numbers)
}

func findTheAnswer(question string) string {
	switch {
	case strings.HasPrefix(question, "which of the following numbers is the largest"):
		return findLargest(question)
	case strings.HasPrefix(question, "what is your name"):
		return "go"
	case strings.HasPrefix(question, "which of the following numbers are primes"):
		return findPrimes(question)
	case strings.HasPrefix(question, "which city is the Eiffel tower in"):
		return "Paris"
	case strings.HasPrefix(question, "what colour is a banana"):
		return "Yellow"
	case strings.HasPrefix(question, "what is the") &&
		strings.HasSuffix(question, "number in the Fibonacci sequence"):
		return findFibonacci(question)
	case strings.HasPrefix(question, "what is"):
		return doSimpleCalculation(question)
	}
	return ""
}

func findLargest(question string) string {
	ints := getIntList(question)

	maxV := 0
	for _, n := range ints {
		maxV = max(n, maxV)
	}
	return strconv.Itoa(maxV)
}

func isPrime(n int) bool {
	for div := 2; div < int(math.Floor(math.Sqrt(float64(n)))); div++ {
		if n%div == 0 {
			return false
		}
	}
	return true
}

func findPrimes(question string) string {
	ints := getIntList(question)

	var strPrimes []string

	for _, n := range ints {
		if isPrime(n) {
			strPrimes = append(strPrimes, strconv.Itoa(n))
		}
	}
	return strings.Join(strPrimes, ", ")
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

func findFibonacci(question string) string {
	re := regexp.MustCompile("\\d+")
	nth, _ := strconv.Atoi(re.FindString(question))

	last, fibo := 0, 1
	for i := 1; i < nth; i++ {
		last, fibo = fibo, fibo+last
	}
	return strconv.Itoa(fibo)
}
