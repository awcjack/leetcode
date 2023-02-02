func copyMap(src map[string]int) map[string]int {
	dst := make(map[string]int, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func matchSubstring(s string, words map[string]int, wordLength int) bool {
	for i := 0; i < len(s); i += wordLength {
		if words[s[i:i+wordLength]] == 0 {
			return false
		} else {
			words[s[i:i+wordLength]]--
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	result := []int{}

	hashMap := make(map[string]int)
	wordLength := len(words[0])
	wordCount := wordLength * len(words)
	for _, word := range words {
		hashMap[word]++
	}

	for i, _ := range s {
		if i+wordCount <= len(s) {
			if matchSubstring(s[i:i+wordCount], copyMap(hashMap), wordLength) {
				result = append(result, i)
			}
		}
	}
	return result
}
