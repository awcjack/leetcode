func uniqueMorseRepresentations(words []string) int {
	if len(words) == 1 {
		return 1
	}

	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	transformedMap := make(map[string]struct{})
	var a byte = 'a'
	for _, word := range words {
		temp := ""
		for i, _ := range word {
			temp += morseCode[word[i]-a]
		}
		transformedMap[temp] = struct{}{}
	}

	return len(transformedMap)
}
