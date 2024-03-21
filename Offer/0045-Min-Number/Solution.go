package leetcode

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	s := make([]string, len(nums))
	for i := range nums {
		s[i] = strconv.Itoa(nums[i])
	}
	quickSort(s, 0, len(s)-1)

	return strings.Join(s, "")
}

// 经典快速排序模板
func quickSort(s []string, left, right int) {
	if left < right {
		p := partition(s, left, right)
		quickSort(s, left, p-1)
		quickSort(s, p+1, right)
	}
}

func partition(s []string, left, right int) int {
	pivot := s[right]
	i, j := left, right

	for i < j {
		for i < j && s[i]+pivot <= pivot+s[i] {
			i++
		}
		for i < j && s[j]+pivot >= pivot+s[j] {
			j--
		}
		s[i], s[j] = s[j], s[i]
	}

	s[i], s[right] = s[right], s[i]
	return i
}
