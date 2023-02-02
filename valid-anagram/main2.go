func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)
	for _, char := range s {
		hashmap[char]++
	}
	for _, char := range t {
		if _, ok := hashmap[char]; !ok {
			return false
		}
		hashmap[char]--
	}
	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}
	return true
}
