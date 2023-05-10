package dp

import "strings"

func CanConstructRec(target string, wordBank []string) bool {
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			if CanConstructRec(remainder, wordBank) {
				return true
			}
		}
	}
	return false
}

func CanConstructDP(target string, wordBank []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}

	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			if CanConstructDP(remainder, wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}
