package leetcode

import "bytes"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := bytes.Buffer{}

	for i := 0; num != 0; {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res.WriteString(symbols[i])
	}

	return res.String()
}
