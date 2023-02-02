func subsetsWithBoundary(nums []int, prefix []int, start int, end int) [][]int {
	sets := [][]int{}
	if start > end {
		return sets
	}
	for i := end; i > start; i-- {
		currentElement := append(prefix, nums[i-1])
		sets = append(sets, currentElement)
		for index := start; index < i-1; index++ {
			currentPrefix := make([]int, len(currentElement)+1)
			copy(currentPrefix, append(currentElement, nums[index]))
			// currentPrefix := append(currentElement, nums[index])[:]
			sets = append(sets, currentPrefix)
			if i-1 > index+1 {
				sets = append(sets, subsetsWithBoundary(nums, currentPrefix, index+1, i-1)...)
			}
		}
	}
	return sets
}

func subsets(nums []int) [][]int {
	sets := [][]int{{}}
	sets = append(sets, subsetsWithBoundary(nums, []int{}, 0, len(nums))...)
	return sets
}
