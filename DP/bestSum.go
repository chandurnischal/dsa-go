package dp

func BestSumRec(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var shortestCombo []int

	for _, num := range nums {
		remainder := target - num
		combo := BestSumRec(remainder, nums)
		if combo != nil {
			combo = append(combo, num)
			if shortestCombo == nil || len(combo) < len(shortestCombo) {
				shortestCombo = combo
			}
		}
	}
	return shortestCombo
}

func BestSumDP(target int, nums []int, memo map[int][]int) []int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var shortestCombo []int

	for _, num := range nums {
		remainder := target - num
		combo := BestSumDP(remainder, nums, memo)
		if combo != nil {
			combo = append(combo, num)
			if shortestCombo == nil || len(combo) < len(shortestCombo) {
				shortestCombo = combo
			}
		}
	}
	memo[target] = shortestCombo
	return memo[target]
}
