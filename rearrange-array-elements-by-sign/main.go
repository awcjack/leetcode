func rearrangeArray(nums []int) []int {
	length := len(nums) / 2
	positiveNums := make([]int, 0, length)
	negativeNums := make([]int, 0, length)
	for _, num := range nums {
		if num >= 0 {
			positiveNums = append(positiveNums, num)
		} else {
			negativeNums = append(negativeNums, num)
		}
	}
	for i, num := range positiveNums {
		nums[i*2] = num
		nums[i*2+1] = negativeNums[i]
	}
	return nums
}
