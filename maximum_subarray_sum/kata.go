package kata

func MaximumSubarraySum(numbers []int) int {
	var max int
	for i, _ := range numbers {
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}