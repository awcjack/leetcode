func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := len(strs[0])
	for i := 1; i < len(strs); i++ {
		length = int(math.Min(float64(length), float64(len(strs[i]))))
	}
	result := ""
	for i := 0; i < length; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return result
			}
		}
		result += string(char)
	}
	return result
}
