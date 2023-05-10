package memoization

import "strings"

func prepend(str string, strs []string) []string {
	res := []string{str}
	res = append(res, strs...)
	return res
}

func Map(word string, words [][]string, prepend func(string, []string) []string) [][]string {
	wordsRes := make([][]string, len(words))

	for i, w := range words {
		wordsRes[i] = prepend(word, w)
	}

	return wordsRes
}

func AllConstructRec(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var res [][]string
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			ways := AllConstructRec(remainder, wordBank)
			ways = Map(word, ways, prepend)
			res = append(res, ways...)
		}
	}

	return res
}

func AllConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return [][]string{{}}
	}
	var res [][]string

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			remainder := target[len(word):]
			ways := AllConstruct(remainder, wordBank, memo)
			ways = Map(word, ways, prepend)
			res = append(res, ways...)
		}
	}

	memo[target] = res
	return res
}
