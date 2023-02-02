func lastRemainingHelper(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 5 {
		return 2
	}
	if n <= 7 {
		return 4
	}
	if n%2 != 0 {
		n = n - 1
	}
	if n%4 == 0 {
		return 6
	}
	return 8
}

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 23 {
		return lastRemainingHelper(n)
	}
	counter := 0
	adjust := make(map[int]bool)
	for {
		if n%2 != 0 {
			n = n - 1
		}
		if (n-2)%4 != 0 {
			adjust[counter] = true
		}
		n = n / 4
		counter++
		if n <= 23 {
			break
		}
	}
	remaining := lastRemainingHelper(n)
	for i := counter; i >= 0; i-- {
		if adjust[i] {
			remaining -= 2
		}
		if i != 0 {
			remaining *= 4
		}

	}
	return remaining
}
