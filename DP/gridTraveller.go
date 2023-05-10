package dp

import "fmt"

func GridTravellerRec(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return GridTravellerRec(m-1, n) + GridTravellerRec(m, n-1)
}

func GridTravellerDP(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	flippedKey := fmt.Sprintf("%d,%d", n, m)
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	if _, ok := memo[flippedKey]; ok {
		return memo[flippedKey]
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = GridTravellerDP(m-1, n, memo) + GridTravellerDP(m, n-1, memo)
	return memo[key]
}
