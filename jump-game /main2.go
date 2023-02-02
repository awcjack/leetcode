// https://leetcode.com/problems/jump-game/
// timeout (change to hashmap)
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	if nums[0] >= len(nums) {
		return true
	}
	for i := nums[0]; i >= 1; i-- {
		if i == len(nums) {
			return true
		}
		if canJump(nums[i:]) {
			return true
		}
	}
	return false
}
