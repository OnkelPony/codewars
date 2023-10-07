package dyn_prog


func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}

	for _, number := range numbers {
		remainder := targetSum - number
		can := canSum(remainder, numbers, memo)
		if can {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}
