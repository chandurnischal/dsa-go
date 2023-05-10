package dp

func CanSumRec(target int, nums []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, num := range nums {
		remainder := target - num
		if CanSumRec(remainder, nums) {
			return true
		}
	}
	return false
}

func CanSumDP(target int, nums []int, memo map[int]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, num := range nums {
		remainder := target - num
		if CanSumDP(remainder, nums, memo) {
			memo[target] = true
			return true
		}
	}
	memo[target] = false
	return false
}
