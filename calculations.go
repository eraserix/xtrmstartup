package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

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

// find answer to questions like 'what is 12 plus 8'. Supports:
// - plus
// - minus
// - multiplied by
// - to the power of
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
