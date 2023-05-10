package memoization

import "strings"

func CountConstructRec(target string, wordBank []string) int {
	if target == "" {
		return 1
	}
	var totalWays int
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			ways := CountConstructRec(remainder, wordBank)
			totalWays += ways
		}
	}
	return totalWays
}

func CountConstruct(target string, wordBank []string, memo map[string]int) int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return 1
	}

	var totalWays int

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			ways := CountConstruct(remainder, wordBank, memo)
			totalWays += ways

		}
	}
	memo[target] = totalWays
	return totalWays
}
