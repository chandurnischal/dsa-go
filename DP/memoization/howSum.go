package dp

func HowSumRec(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, num := range nums {
		remainder := target - num
		res := HowSumRec(remainder, nums)
		if res != nil {
			res = append(res, num)
			return res
		}
	}
	return nil
}

func HowSumDP(target int, nums []int, memo map[int][]int) []int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, num := range nums {
		remainder := target - num
		res := HowSumDP(remainder, nums, memo)
		if res != nil {
			memo[target] = append(res, num)
			return memo[target]
		}
	}
	memo[target] = nil
	return memo[target]
}
