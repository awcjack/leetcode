func findDuplicate(nums []int) int {
	numsHashmap := make([]bool, len(nums))
	for _, num := range nums {
		if numsHashmap[num] != false {
			return num
		}
		numsHashmap[num] = true
	}
	return 0
}
