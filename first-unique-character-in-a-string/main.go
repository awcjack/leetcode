func firstUniqChar(s string) int {
	hashmap := make(map[byte][]int)
	for i, _ := range s {
		hashmap[s[i]] = append(hashmap[s[i]], i)
	}
	result := -1
	for _, intArray := range hashmap {
		if len(intArray) == 1 {
			if result == -1 {
				result = intArray[0]
			} else if intArray[0] < result {
				result = intArray[0]
			}
		}
	}
	return result
}
