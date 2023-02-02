// out of memory

func isConnected(origin string, destination string) bool {
	diff := 0
	for i, v := range origin {
		if byte(v) != destination[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func buildAdjacentList(wordList []string) map[string]map[string]struct{} {
	adjacentList := make(map[string]map[string]struct{})

	for _, word := range wordList {
		adjacentList[word] = make(map[string]struct{})
		for _, otherWord := range wordList {
			if isConnected(word, otherWord) {
				adjacentList[word][otherWord] = struct{}{}
			}
		}
	}

	return adjacentList
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	found := false
	for _, v := range wordList {
		if v == endWord {
			found = true
		}
	}
	if !found {
		return [][]string{}
	}

	wordList = append(wordList, beginWord)
	adjacentList := buildAdjacentList(wordList)

	result := [][]string{}
	layer := []string{beginWord}
	used := make(map[string]struct{})
	found = false
	for !found {
		temp := make([]string, 0, len(layer))
		innerUsed := make(map[string]struct{})
		for _, node := range layer {
			cursorArray := strings.Split(node, ".")
			cursor := cursorArray[len(cursorArray)-1]
			for k, _ := range adjacentList[cursor] {
				if _, ok := used[k]; !ok {
					if k == endWord {
						found = true
						tempResult := append(cursorArray, k)
						result = append(result, tempResult)
					} else {
						temp = append(temp, fmt.Sprintf("%s.%s", node, k))
						innerUsed[k] = struct{}{}
					}
				}
			}
		}
		if len(innerUsed) == 0 && !found {
			return [][]string{}
		}
		for k, _ := range innerUsed {
			used[k] = struct{}{}
		}
		layer = temp
	}
	return result
}
