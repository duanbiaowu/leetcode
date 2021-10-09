package leetcode

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return digits
	}
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
