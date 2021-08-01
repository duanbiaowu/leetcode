package leetcode

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}

	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := div(dividend, divisor)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return res
	}

	return -res
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	sum := b
	for sum+sum <= a {
		count += count
		sum += sum
	}
	return count + div(a-sum, b)
}
