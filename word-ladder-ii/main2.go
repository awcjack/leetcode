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
	alternative := make(map[string][]string)
	found = false
	for !found {
		temp := make([]string, 0, len(layer))
		innerUsed := make(map[string][]string)
		for _, node := range layer {
			cursor := node[len(node)-len(beginWord):]
			for k, _ := range adjacentList[cursor] {
				if _, ok := used[k]; !ok {
					if _, ok := innerUsed[k]; ok {
						innerUsed[k] = append(innerUsed[k], node+"."+k)
						for _, v := range alternative[node] {
							alternative[innerUsed[k][0]] = append(alternative[innerUsed[k][0]], v+"."+k)
						}
					} else if k == endWord {
						found = true
						cursorArray := strings.Split(node, ".")
						tempResult := append(cursorArray, k)
						result = append(result, tempResult)
						if alternatives, ok := alternative[node+"."+k]; ok {
							for _, v := range alternatives {
								cursorArray := strings.Split(v+node[len(v):], ".")
								tempResult := append(cursorArray, k)
								result = append(result, tempResult)
							}
						}
					} else {
						temp = append(temp, node+"."+k)
						innerUsed[k] = append(innerUsed[k], node+"."+k)
						if alternatives, ok := alternative[node]; ok {
							alternative[node+"."+k] = append(alternative[node+"."+k], alternatives...)
						}
					}
				}
			}
			delete(alternative, node)
		}
		if len(innerUsed) == 0 && !found {
			return [][]string{}
		}
		for k, v := range innerUsed {
			used[k] = struct{}{}
			if len(v) > 1 {
				alternative[v[0]] = append(alternative[v[0]], v[1:]...)
			}
			delete(innerUsed, k)
		}
		layer = temp
	}
	return result
}
