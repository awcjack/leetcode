func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	maxLength := len(num1) + len(num2)
	result := make([]byte, maxLength)
	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		for j := len(num2) - 1; j >= 0; j-- {
			sum := int(num1[i]-'0')*int(num2[j]-'0') + carry
			if result[maxLength-(len(num1)-1-i)-(len(num2)-1-j)-1] > 0 {
				sum += int(result[maxLength-(len(num1)-1-i)-(len(num2)-1-j)-1]) - int('0')
			}
			result[maxLength-(len(num1)-1-i)-(len(num2)-1-j)-1] = byte(sum%10) + '0'
			carry = sum / 10
		}
		if carry != 0 {
			result[maxLength-(len(num1)-1-i)-len(num2)-1] += byte(carry) + '0'
		}

	}
	if result[0] == 0 {
		return string(result[1:])
	}
	return string(result[:])
}
