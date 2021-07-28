package leetcode

func firstBadVersion(n int) int {
	low, hi := 1, n
	for low < hi {
		mid := low + (hi-low)>>1
		if isBadVersion(mid) {
			hi = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// mock error version for testcases
func isBadVersion(version int) bool {
	if version >= 3 {
		return true
	}
	return false
}
