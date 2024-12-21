package leetcode

import (
	"math/rand"
)

// 第K个最大元素: len(nums)-k
// 第K个最小元素: k
// 快速排序
// 问题: 遇到极端边界测试用例时会超时
// https://leetcode.cn/problems/kth-largest-element-in-an-array/solutions/307351/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcod-2/comments/2352693
func findKthLargest(nums []int, k int) int {
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

// 基于快速排序思路存在的问题
// 改造过后的版本
func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickSelectOpt(nums, 0, n-1, n-k)
}

func quickSelectOpt(nums []int, left, right, k int) int {
	if left == right {
		return nums[k]
	}

	pivot := nums[left]

	i, j := left-1, right+1
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		return quickSelect(nums, left, j, k)
	}
	return quickSelect(nums, j+1, right, k)
}

// 分治法
func findKthLargestDivide(nums []int, k int) int {
	// 直接以最右边元素作为基准值
	n := len(nums)
	pivot := nums[n-1]

	// 将大于等于小于的元素分别放入三个数组
	small, equal, large := []int{}, []int{}, []int{}

	for _, num := range nums {
		if num < pivot {
			small = append(small, num)
		} else if num > pivot {
			large = append(large, num)
		} else {
			equal = append(equal, num)
		}
	}

	// 第 K 个最大元素在大的数组中
	if k <= len(large) {
		return findKthLargestDivide(large, k)
	}

	// 第 K 个最大元素在小的数组中
	if len(large)+len(equal) < k {
		return findKthLargestDivide(small, k-len(large)-len(equal))
	}

	// 第 K 个最大元素在相等的数组中
	// 无论相等的数组中有多少个元素，都不影响结果
	// 直接返回基准元素即可
	return pivot
}

// 首先建立一个大顶堆 (根元素为数组中的最大元素)
// 然后执行 k−1 次删除操作
// 堆顶元素就是数组中的第K个最大元素
// 数组的具体元素变化过程可以参考官方题解:
// https://leetcode.cn/problems/kth-largest-element-in-an-array/solutions/307351/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcod-2
func findKthLargestWithHeap(nums []int, k int) int {
	n := len(nums)
	heapSize := n
	buildHeap(nums, heapSize)

	for i := n - 1; i >= n-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}

	return nums[0]
}

func buildHeap(nums []int, size int) {
	for i := size>>1 - 1; i >= 0; i-- {
		maxHeapify(nums, i, size)
	}
}

func maxHeapify(nums []int, index, size int) {
	left, right, largest := index*2+1, index*2+2, index

	if left < size && nums[left] > nums[largest] {
		largest = left
	}
	if right < size && nums[right] > nums[largest] {
		largest = right
	}
	if largest != index {
		nums[index], nums[largest] = nums[largest], nums[index]
		maxHeapify(nums, largest, size)
	}
}
