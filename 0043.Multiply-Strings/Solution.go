package leetcode

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	arr := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j]) - '0'
			arr[i+j+1] += x * y
		}
	}

	for i := m + n - 1; i > 0; i-- {
		arr[i-1] += arr[i] / 10
		arr[i] %= 10
	}

	var res strings.Builder
	index := 0
	if arr[0] == 0 {
		index = 1
	}
	for index < m+n {
		res.WriteString(strconv.Itoa(arr[index]))
		index++
	}
	return res.String()
}
