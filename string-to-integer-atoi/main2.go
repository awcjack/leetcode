var MAX int = int(math.Pow(2, 31)) - 1
var MIN int = -int(math.Pow(2, 31))

func myAtoi(s string) int {
	trimmedString := strings.TrimSpace(s)
	negative := false
	symbol := false
	if len(trimmedString) == 0 {
		return 0
	}
	if trimmedString[0]-'-' == 0 {
		negative = true
		symbol = true
	} else if trimmedString[0]-'+' == 0 {
		symbol = true
	}
	if symbol {
		trimmedString = trimmedString[1:]
	}
	n := make([]byte, 0, len(trimmedString))
	for _, r := range trimmedString {
		if r >= '0' && r <= '9' {
			n = append(n, byte(r))
		} else {
			break
		}
	}
	trimmedString = string(n)
	trimmedString = strings.TrimLeft(trimmedString, "0")
	var number int = 0
	if !negative && len(trimmedString) > 10 {
		return MAX
	} else if negative && len(trimmedString) > 10 {
		return MIN
	}
	for i, r := range trimmedString {
		if r >= '0' && r <= '9' {
			number += int(r-'0') * int(math.Pow(10, float64((len(trimmedString)-i-1))))
		} else {
			break
		}
	}
	if !negative && number > MAX {
		return MAX
	} else if negative && number > MAX+1 {
		return MIN
	}
	if negative {
		return -number
	}
	return number
}
