func isPossible(nums []int) bool {
	numsHashmap := make(map[int]int)
	sequenceHashmap := make(map[int]int)

	for _, v := range nums {
		numsHashmap[v]++
	}

	for _, v := range nums {
		if numsHashmap[v] != 0 {
			if sequenceHashmap[v-1] > 0 {
				numsHashmap[v]--
				sequenceHashmap[v-1]--
				sequenceHashmap[v]++
			} else {
				if numsHashmap[v+1] == 0 || numsHashmap[v+2] == 0 {
					return false
				}
				numsHashmap[v]--
				numsHashmap[v+1]--
				numsHashmap[v+2]--
				sequenceHashmap[v+2]++
			}
		}
	}
	return true
}
