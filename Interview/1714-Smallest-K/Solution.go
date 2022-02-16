package leetcode

import (
	"math/rand"
	"time"
)

func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k <= 0 {
		return nil
	}
	// 引入随机 base 避免时间复杂度退化为 N^2
	rand.Seed(time.Now().UnixNano())
	quickSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSort(nums []int, left, right, k int) {
	if left < right {
		index := partition(nums, left, right)
		if k < index {
			quickSort(nums, left, index-1, k)
		} else if k > index {
			quickSort(nums, index+1, right, k)
		}
	}
}

func partition(nums []int, left, right int) int {
	// 引入随机 base 元素
	p := left + rand.Intn(right-left+1)
	nums[p], nums[right] = nums[right], nums[p]

	pivot := nums[right]
	pos := left
	for left < right {
		if nums[left] < pivot {
			nums[pos], nums[left] = nums[left], nums[pos]
			pos++
		}
		left++
	}
	nums[pos], nums[right] = nums[right], nums[pos]
	return pos
}
