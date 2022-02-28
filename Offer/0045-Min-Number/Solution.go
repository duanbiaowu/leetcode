package leetcode

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	quickSort(strs, 0, len(strs)-1)

	var sb strings.Builder
	for i := range strs {
		sb.WriteString(strs[i])
	}
	return sb.String()
}

func quickSort(strs []string, left, right int) {
	if left >= right {
		return
	}

	pivot := strs[right]
	i, j := left, right
	for i < j {
		for i < j && strs[i]+pivot <= pivot+strs[i] {
			i++
		}
		for i < j && strs[j]+pivot >= pivot+strs[j] {
			j--
		}
		strs[i], strs[j] = strs[j], strs[i]
	}
	strs[i], strs[right] = strs[right], strs[i]

	quickSort(strs, left, i-1)
	quickSort(strs, i+1, right)
}
