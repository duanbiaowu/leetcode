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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n&1 == 0 {
		mid := n >> 1
		return float64(getKthElement(nums1, nums2, mid)+getKthElement(nums1, nums2, mid+1)) / 2
	}
	return float64(getKthElement(nums1, nums2, n>>1+1))
}

func getKthElement(nums1, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	if k == 0 || (m == 0 && n == 0) {
		return 0
	}

	i, j := 0, 0
	for {
		if i == m {
			return nums2[j+k-1]
		}
		if j == n {
			return nums1[i+k-1]
		}
		if k == 1 {
			return min(nums1[i], nums2[j])
		}

		mid := k >> 1
		newI := min(i+mid, m) - 1
		newJ := min(j+mid, n) - 1

		if nums1[newI] <= nums2[newJ] {
			k -= newI - i + 1
			i = newI + 1
		} else {
			k -= newJ - j + 1
			j = newJ + 1
		}
	}
}

// 划分数组
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays2(nums2, nums1)
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
		numsIm1 := math.MinInt32
		if i != 0 {
			numsIm1 = nums1[i-1]
		}
		numsI := math.MaxInt32
		if i != m {
			numsI = nums1[i]
		}
		numsJm1 := math.MinInt32
		if j != 0 {
			numsJm1 = nums2[j-1]
		}
		numsJ := math.MaxInt32
		if j != n {
			numsJ = nums2[j]
		}

		if numsIm1 <= numsJ {
			median1 = max(numsIm1, numsJm1)
			median2 = min(numsI, numsJ)
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
