package leetcode

import (
	"strings"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res strings.Builder

	for i := 0; num != 0; num -= values[i] {
		for values[i] > num {
			i++
		}
		res.WriteString(symbols[i])
	}

	return res.String()
}
