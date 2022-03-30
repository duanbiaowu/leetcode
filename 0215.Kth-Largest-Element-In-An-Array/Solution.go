package leetcode

import (
	"math/rand"
	"time"
)

// 第K个最大元素: len(nums)-k
// 第K个最小元素: k
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, index int) int {
	pivot := randomPartition(nums, left, right)
	if pivot == index {
		return nums[pivot]
	} else if pivot < index {
		return quickSelect(nums, pivot+1, right, index)
	}
	return quickSelect(nums, left, pivot-1, index)
}

func randomPartition(nums []int, left, right int) int {
	index := rand.Int()%(right-left+1) + left
	nums[index], nums[right] = nums[right], nums[index]
	return partition(nums, left, right)
}

func partition(nums []int, left, right int) int {
	pivot := right
	for left < right {
		for left < right && nums[left] <= nums[pivot] {
			left++
		}
		for left < right && nums[right] >= nums[pivot] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]
	return left
}
