package tabulation

func FibReg(n int) int {
	if n <= 1 {
		return n
	}

	prev := 0
	current := 1

	for i := 2; i <= n; i++ {
		next := prev + current
		prev = current
		current = next
	}

	return current
}

func Fib(n int) int {
	table := make([]int, n+1)

	table[1] = 1

	for i := 0; i <= n-2; i++ {
		table[i+1] += table[i]
		table[i+2] += table[i]
	}
	return table[n]
}
