func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if 'a' <= b && b <= 'z' || '0' <= b && b <= '9' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func isPalindrome(s string) bool {
	lowerS := strings.ToLower(s)
	cleanStr := clean([]byte(lowerS))
	for i := 0; i < len(cleanStr)/2; i++ {
		if cleanStr[i] != cleanStr[len(cleanStr)-i-1] {
			return false
		}
	}
	return true
}
