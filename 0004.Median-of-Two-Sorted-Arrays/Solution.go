package leetcode

// 暴力解法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	left, right, cur := 0, 0, 0
	n1, n2 := len(nums1), len(nums2)
	newArray := make([]int, n1 + n2)

	for left < n1 && right < n2 {
		if nums1[left] < nums2[right] {
			newArray[cur] = nums1[left]
			left++
		} else {
			newArray[cur] = nums2[right]
			right++
		}
		cur++
	}

	for left < n1 {
		newArray[cur] = nums1[left]
		cur++
		left++
	}
	for right < n2 {
		newArray[cur] = nums2[right]
		cur++
		right++
	}

	n := len(newArray)
	if n == 0 {
		return 0
	} else if n&1 == 1 {
		return float64(newArray[n/2])
	} else {
		return float64(newArray[n/2] + newArray[n/2-1]) / 2
	}
}
