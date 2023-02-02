func generateParenthesisHelper(bracket int, n int) []string {
	var result = []string{}

	if n == 1 {
		return []string{"()"}
	}

	if bracket == 1 {
		for i := 0; i < n-1; i++ {
			inside := generateParenthesisHelper(i+1, n-1)
			for _, v := range inside {
				result = append(result, fmt.Sprintf("(%s)", v))
			}
		}
		return result
	}

	for j := n + 1 - bracket; j > 0; j-- {
		leftmosts := generateParenthesisHelper(1, j)
		remainings := generateParenthesisHelper(bracket-1, n-j)
		for _, remaining := range remainings {
			for _, leftmost := range leftmosts {
				result = append(result, fmt.Sprintf("%s%s", leftmost, remaining))
			}
		}
	}
	return result
}

func generateParenthesis(n int) []string {
	var result = []string{}
	for i := 0; i < n; i++ {
		current := generateParenthesisHelper(i+1, n)

		result = append(result, current...)
	}
	return result
}
