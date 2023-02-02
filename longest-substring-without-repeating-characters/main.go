func lengthOfLongestSubstring(s string) int {
	charHashmap := make(map[byte]int, 26)
	max := 0
	length := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if charHashmap[s[i]] > start {
			if length > max {
				max = length
			}
			length -= charHashmap[s[i]] - start
			if length < 0 {
				length = 0
			}
			start = charHashmap[s[i]]
		}
		charHashmap[s[i]] = i + 1
		length++
	}
	if length > max {
		max = length
	}
	return max
}
