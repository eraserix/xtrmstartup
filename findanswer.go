package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getIntList(question string) []int {
	numbers := getStringList(question)
	ints := make([]int, len(numbers))
	for i, n := range numbers {
		ints[i], _ = strconv.Atoi(n)
	}
	return ints
}

func getStringList(question string) []string {
	stringlist := strings.Split(strings.Split(question, ":")[1], ",")
	result := make([]string, len(stringlist))
	for i, s := range stringlist {
		result[i] = strings.TrimSpace(s)
	}
	return result
}

func findTheAnswer(question string) string {
	switch {
	case strings.HasPrefix(question, "which of the following numbers is the largest"):
		return findLargest(question)
	case strings.HasPrefix(question, "what is your name"):
		return "go"
	case strings.HasPrefix(question, "which of the following numbers are primes"):
		return findPrimes(question)
	case strings.HasPrefix(question, "which of the following numbers is both a square and a cube"):
		return findSquareAndCube(question)
	case strings.HasPrefix(question, "which city is the Eiffel tower in"):
		return "Paris"
	case strings.HasPrefix(question, "what colour is a banana"):
		return "Yellow"
	case strings.HasPrefix(question, "what is the english scrabble score of"):
		return findScrabbleScore(question)
	case strings.HasPrefix(question, "which of the following is an anagram of"):
		return findAnagram(question)
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

func doPlusQuestion(question string) string {
	var a, b, c int
	switch {
	case strings.Count(question, "plus") == 2:
		fmt.Sscanf(question, "what is %d plus %d plus %d", &a, &b, &c)
		return strconv.Itoa(a + b + c)
	case strings.Count(question, "plus") == 1 &&
		strings.Count(question, "multiplied by") == 1:
		_, e := fmt.Sscanf(question, "what is %d plus %d multiplied by %d", &a, &b, &c)
		if e == nil {
			return strconv.Itoa(a + b*c)
		}
		_, e = fmt.Sscanf(question, "what is %d multiplied by %d plus %d", &a, &b, &c)
		if e == nil {
			return strconv.Itoa(a*b + c)
		}
	case strings.Count(question, "plus") == 1:
		fmt.Sscanf(question, "what is %d plus %d", &a, &b)
		return strconv.Itoa(a + b)
	}
	return ""
}

func doSimpleCalculation(question string) string {
	var a, b int
	switch {
	case strings.Contains(question, "plus"):
		return doPlusQuestion(question)
	case strings.Contains(question, "minus"):
		fmt.Sscanf(question, "what is %d minus %d", &a, &b)
		return strconv.Itoa(a - b)
	case strings.Contains(question, "multiplied by"):
		fmt.Sscanf(question, "what is %d multiplied by %d", &a, &b)
		return strconv.Itoa(a * b)
	case strings.Contains(question, "to the power of"):
		fmt.Sscanf(question, "what is %d to the power of %d", &a, &b)
		base := new(big.Int).SetInt64(int64(a))
		result := new(big.Int).Set(base)
		for i := 1; i < b; i++ {
			result.Mul(result, base)
		}
		return result.String()
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

func findSquareAndCube(question string) string {
	isIntPower := func(number, power int) bool {
		root := math.Pow(float64(number), 1/float64(power))
		num2 := int(math.Pow(math.Floor(root+0.5), float64(power)))
		return num2 == number
	}

	ints := getIntList(question)
	var result []string

	for _, n := range ints {
		if isIntPower(n, 2) && isIntPower(n, 3) {
			result = append(result, strconv.Itoa(n))
		}
	}
	return strings.Join(result, ", ")
}

func findScrabbleScore(question string) string {
	var word string
	fmt.Sscanf(question, "what is the english scrabble score of %s", &word)

	score := 0

	// 1 point: E ×12, A ×9, I ×9, O ×8, N ×6, R ×6, T ×6, L ×4, S ×4, U ×4
	// 2 points: D ×4, G ×3
	// 3 points: B ×2, C ×2, M ×2, P ×2
	// 4 points: F ×2, H ×2, V ×2, W ×2, Y ×2
	// 5 points: K ×1
	// 8 points: J ×1, X ×1
	// 10 points: Q ×1, Z ×1
	for _, char := range word {
		switch char {
		case 'e', 'a', 'i', 'o', 'n', 'r', 't', 'l', 's', 'u':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	return strconv.Itoa(score)
}

func findAnagram(question string) string {
	var word string
	fmt.Sscanf(question, "which of the following is an anagram of \"%s\":", &word)
	word = word[:len(word)-2]
	runes := strings.Split(word, "")
	sort.Strings(runes)
	candidates := getStringList(question)

	var result []string
	for _, c := range candidates {
		runesC := strings.Split(c, "")
		sort.Strings(runesC)
		if reflect.DeepEqual(runesC, runes) {
			result = append(result, c)
		}
	}

	return strings.Join(result, ", ")
}
