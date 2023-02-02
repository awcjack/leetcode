

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	newStrs := make([]string, len(strs))
	for i, str := range strs {
		runeArray := []rune(str)
		sort.Sort(sortRuneString(runeArray))
		newStrs[i] = string(runeArray)
	}
	hashmap := make(map[string][]int)
	for i, str := range newStrs {
		hashmap[str] = append(hashmap[str], i)
	}
	result := make([][]string, 0, len(hashmap))
	for _, val := range hashmap {
		subResult := make([]string, 0, len(val))
		for _, value := range val {
			subResult = append(subResult, strs[value])
		}
		result = append(result, subResult)
	}
	return result
}
