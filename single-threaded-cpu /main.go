type minHeap [][]int

func (h *minHeap) Push(num interface{}) {
	*h = append(*h, num.([]int))
}

func (h *minHeap) Pop() interface{} {
	length := len(*h)
	result := (*h)[length-1]
	*h = (*h)[:length-1]
	return result
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][2] < h[j][2]
	}
	return h[i][1] < h[j][1]
}

func getOrder(tasks [][]int) []int {
	sortedTasks := make([][]int, len(tasks))
	for i, task := range tasks {
		sortedTasks[i] = make([]int, 3)
		sortedTasks[i][0] = task[0]
		sortedTasks[i][1] = task[1]
		sortedTasks[i][2] = i
	}

	sort.Slice(sortedTasks, func(i, j int) bool {
		return sortedTasks[i][0] < sortedTasks[j][0]
	})

	currentTime := sortedTasks[0][0]
	result := make([]int, 0, len(tasks))

	index := 0
	h := &minHeap{}

	for index < len(tasks) || h.Len() > 0 {
		if index < len(tasks) && sortedTasks[index][0] <= currentTime {
			heap.Push(h, sortedTasks[index])
			index++
		} else if h.Len() > 0 {
			task := heap.Pop(h).([]int)
			currentTime += task[1]
			result = append(result, task[2])
		} else {
			currentTime = sortedTasks[index][0]
		}
	}

	return result
}
