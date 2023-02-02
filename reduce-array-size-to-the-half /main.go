func minSetSize(arr []int) int {
	heatmap := make(map[int]int)
	for _, v := range arr {
		heatmap[v]++
	}

	keys := make([]int, 0, len(heatmap))

	for key := range heatmap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return heatmap[keys[i]] > heatmap[keys[j]]
	})

	var half int = (len(arr) + 1) / 2
	counter := 0
	for i, v := range keys {
		counter += heatmap[v]
		if counter >= half {
			return i + 1
		}
	}
	return 1
}
