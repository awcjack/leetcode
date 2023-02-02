func reverseWords(s string) string {
	lastSpaceIndex := -1
	stringByte := []byte(s)
	length := len(stringByte)
	for i := 0; i <= length; i++ {
		// byte of space character is 32
		if i == length || stringByte[i] == 32 {
			startIndex := lastSpaceIndex + 1
			endIndex := i - 1
			for startIndex < endIndex {
				temp := stringByte[startIndex]
				stringByte[startIndex] = stringByte[endIndex]
				stringByte[endIndex] = temp
				startIndex++
				endIndex--
			}
			lastSpaceIndex = i
		}
	}
	return string(stringByte)
}
