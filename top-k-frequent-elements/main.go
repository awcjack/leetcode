func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)

	for _, val := range nums {
		hashmap[val]++
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv
	for k, v := range hashmap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, ss[i].Key)
	}
	return result
}
