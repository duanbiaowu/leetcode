package leetcode

func addStrings(num1 string, num2 string) string {
	left, right := len(num1)-1, len(num2)-1
	var buf []byte
	cur := 0

	for left >= 0 || right >= 0 {
		if left >= 0 {
			cur += int(num1[left] - '0')
			left--
		}
		if right >= 0 {
			cur += int(num2[right] - '0')
			right--
		}
		buf = append(buf, byte(cur%10+'0'))
		cur /= 10
	}

	left, right = 0, len(buf)-1
	for left < right {
		buf[left], buf[right] = buf[right], buf[left]
		left++
		right--
	}

	return string(buf)
}
