func findNthDigit(n int) int {
	digitMinusOne := int(math.Log10(float64(n)))
	if digitMinusOne == 0 {
		return n
	}
	counter := 0
	for i := 0; i < digitMinusOne; i++ {
		if n >= 9*int(math.Pow10(i))*(i+1) {
			n -= 9 * int(math.Pow10(i)) * (i + 1)
			counter++
		}
	}
	if n == 0 {
		return 9
	}
	if n%(counter+1) == 0 {
		if n/(counter+1)%10 == 0 {
			return 9
		}
		return (n/(counter+1))%10 - 1
	}
	number := n / (counter + 1)
	digit := (counter + 1) - (n % (counter + 1))
	number += int(math.Pow10((counter)))
	return (number / int(math.Pow10(digit))) % 10
}
