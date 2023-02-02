func combinationSum(candidates []int, target int) [][]int {
	hashmap := make(map[int][][]int)
	sort.Ints(candidates)
	for _, v := range candidates {
		hashmap[v] = [][]int{{v}}
	}
	for _, v := range candidates {
		for i := 0; i <= target; i++ {
			element := hashmap[i]
			if len(element) != 0 {
				for _, array := range element {
					if v+i > target {
						break
					}
					if array[len(array)-1] > v {
						continue
					}
					temp := make([]int, len(array)+1)
					temp2 := append(array, v)
					copy(temp, temp2)
					temp2 = nil
					// sort.Ints(temp)
					hashmap[v+i] = append(hashmap[v+i], temp)
					temp = nil
				}
			}
			// fmt.Println("element", element, "hashmap", hashmap)
		}
		// fmt.Println("v", v, "hashmap", hashmap)
	}
	return hashmap[target]
}
