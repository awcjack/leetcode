func isValid(s string) bool {
	stack := make([]string, 0, (len(s)+1)/2)
	for _, char := range s {
		if string(char) == "(" || string(char) == "{" || string(char) == "[" {
			stack = append(stack, string(char))
		} else {
			if len(stack) == 0 {
				return false
			}
			if string(char) == ")" && stack[len(stack)-1] != "(" {
				return false
			} else if string(char) == "]" && stack[len(stack)-1] != "[" {
				return false
			} else if string(char) == "}" && stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
