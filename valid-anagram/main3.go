func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// not supporting unicode char
	hashmap := make(map[byte]int)
	for i, _ := range s {
		hashmap[s[i]]++
		hashmap[t[i]]--
	}
	for _, val := range hashmap {
		if val != 0 {
			return false
		}
	}
	return true
}
