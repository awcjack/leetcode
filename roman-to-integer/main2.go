func romanToInt(s string) int {
	keyNumberMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	hashmap := make(map[string]int)
	for i := len(s) - 1; i >= 0; i-- {
		currentChar := string(s[i])
		if i != 0 {
			prevChar := string(s[i-1])
			if currentChar == "V" || currentChar == "X" {
				if prevChar == "I" {
					result += keyNumberMap[currentChar] - keyNumberMap[prevChar]
					i--
					continue
				}
			} else if currentChar == "L" || currentChar == "C" {
				if prevChar == "X" {
					result += keyNumberMap[currentChar] - keyNumberMap[prevChar]
					i--
					continue
				}
			} else if currentChar == "D" || currentChar == "M" {
				if prevChar == "C" {
					result += keyNumberMap[currentChar] - keyNumberMap[prevChar]
					i--
					continue
				}
			}
			if prevChar != currentChar {
				temp := hashmap[currentChar] + 1
				hashmap[currentChar] = 0
				result += temp * keyNumberMap[currentChar]
			} else {
				hashmap[currentChar]++
			}
		} else {
			temp := hashmap[currentChar] + 1
			hashmap[currentChar] = 0
			result += temp * keyNumberMap[currentChar]
		}

	}

	return result
}
