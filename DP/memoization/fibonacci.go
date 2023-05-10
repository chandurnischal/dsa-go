package memoization

func FibonacciRec(n int) int {
	if n <= 2 {
		return 1
	}
	return FibonacciRec(n-1) + FibonacciRec(n-2)
}

// returns nth Fibonacci number
func Fibonacci(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
	return memo[n]
}
