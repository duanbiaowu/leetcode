package leetcode

import (
	"math"
)

// 暴力解法
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	left, right, cur := 0, 0, 0
//	n1, n2 := len(nums1), len(nums2)
//	newArray := make([]int, n1 + n2)
//
//	for left < n1 && right < n2 {
//		if nums1[left] < nums2[right] {
//			newArray[cur] = nums1[left]
//			left++
//		} else {
//			newArray[cur] = nums2[right]
//			right++
//		}
//		cur++
//	}
//
//	for left < n1 {
//		newArray[cur] = nums1[left]
//		cur++
//		left++
//	}
//	for right < n2 {
//		newArray[cur] = nums2[right]
//		cur++
//		right++
//	}
//
//	n := n1 + n2
//	if n == 0 {
//		return 0
//	} else if n&1 == 1 {
//		return float64(newArray[n/2])
//	} else {
//		return float64(newArray[n/2] + newArray[n/2-1]) / 2.0
//	}
//}

// 二分查找
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n := len(nums1) + len(nums2)
//	if n&1 == 0 {
//		mid1, mid2 := n>>1-1, n>>1
//		return float64(getKthElement(nums1, nums2, mid1+1)+getKthElement(nums1, nums2, mid2+1)) / 2.0
//	} else {
//		return float64(getKthElement(nums1, nums2, n>>1+1))
//	}
//}

func getKthElement(nums1, nums2 []int, k int) int {
	if k == 0 || (len(nums1) == 0 && len(nums2) == 0) {
		return 0
	}

	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		mid := k >> 1
		newIndex1 := min(index1+mid, len(nums1)) - 1
		newIndex2 := min(index2+mid, len(nums2)) - 1

		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

// 划分数组
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m
	// median1：前一部分的最大值
	// median2：后一部分的最小值
	median1, median2 := 0, 0

	for left <= right {
		// 前一部分包含 nums1[0 .. i-1] 和 nums2[0 .. j-1]
		// 后一部分包含 nums1[i .. m-1] 和 nums2[j .. n-1]
		i := (left + right) >> 1
		j := (m+n+1)>>1 - i
		// nums_im1 表示 nums1[i-1]
		// nums_i 表示 nums1[i]
		// nums_jm1 表示 nums2[j-1]
		// nums_j 表示 nums2[j]
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}

		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)&1 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
