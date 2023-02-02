// cheating
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
			fmt.Println(string(r))
			break
		}
	}
	trimmedString = string(n)
	number, _ := strconv.Atoi(trimmedString)
	if !negative && number > int(math.Pow(2, 31))-1 {
		return int(math.Pow(2, 31)) - 1
	} else if negative && number > int(math.Pow(2, 31)) {
		return -int(math.Pow(2, 31))
	}
	if negative {
		return -number
	}
	return number
}
