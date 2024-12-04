package leetcode

import "sort"

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}

	sort.Ints(citations)
	low, hi := 0, n-1
	for low < hi {
		mid := low + (hi-low)>>1
		if citations[mid] < n-mid {
			low = mid + 1
		} else {
			hi = mid
		}
	}

	// 左边界
	if citations[low] == 0 {
		return 0
	}
	return n - low
}

func hIndex2(citations []int) int {
	sort.Ints(citations)

	h := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}

	return h
}
