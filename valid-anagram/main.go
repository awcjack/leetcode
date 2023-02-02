func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)
	for _, char := range s {
		hashmap[char]++
	}
	for _, char := range t {
		hashmap[char]--
	}
	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}
	return true
}
