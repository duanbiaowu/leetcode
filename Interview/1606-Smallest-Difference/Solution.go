package leetcode

import (
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	res := math.MaxInt32
	for index := range a {
		low, hi := 0, len(b)-1
		for low <= hi {
			mid := low + (hi-low)>>1
			diff := a[index] - b[mid]
			res = min(res, abs(diff))
			if diff == 0 {
				return res
			} else if diff > 0 {
				low = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
