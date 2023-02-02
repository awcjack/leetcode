func twoSum(numbers []int, target int) []int {
	for i := 1; i <= len(numbers); i++ {
		for j := i + 1; j <= len(numbers); j++ {
			if numbers[i-1]+numbers[j-1] == target {
				return []int{i, j}
			}
			if numbers[i-1]+numbers[j-1] > target {
				break
			}
		}
	}
	return []int{0, 0}
}
