func canJump(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	if nums[0] >= len(nums)-1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	threshold := 0
	for i := len(nums) - 2; i >= 0; i-- {
		threshold++
		if nums[i] >= threshold {
			threshold = 0
		}
	}
	if nums[0] >= threshold {
		return true
	}

	return false
}
