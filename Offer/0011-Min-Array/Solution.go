package leetcode

func minArray(numbers []int) int {
	low, hi := 0, len(numbers)-1
	if hi < 0 {
		return -1
	}
	// 理解此处条件和标准二分查找条件 (low <= hi) 的区别
	for low < hi {
		mid := low + (hi-low)>>1
		if numbers[mid] < numbers[hi] {
			hi = mid
		} else if numbers[mid] > numbers[hi] {
			low = mid + 1
		} else {
			hi--
		}
	}
	return numbers[low]
}
