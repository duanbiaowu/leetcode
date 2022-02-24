package leetcode

// reference: https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/solution/jian-zhi-offer-51-shu-zu-zhong-de-ni-xu-pvn2h/
func reversePairs(nums []int) int {
	cnt := 0
	merge(nums, 0, len(nums)-1, &cnt)
	return cnt
}

func merge(nums []int, low, hi int, cnt *int) {
	if low < hi {
		mid := low + (hi-low)>>1
		merge(nums, low, mid, cnt)
		merge(nums, mid+1, hi, cnt)
		mergeSort(nums, low, mid, hi, cnt)
	}
}

// 合并阶段 本质上是 合并两个排序数组 的过程，
// 而每当遇到 左子数组当前元素 > 右子数组当前元素 时，
// 意味着 「左子数组当前元素 至 末尾元素」 与 「右子数组当前元素」 构成了若干 「逆序对」
func mergeSort(nums []int, low, mid, hi int, cnt *int) {
	tmp := make([]int, hi-low+1)
	index := 0
	p1, p2 := low, mid+1

	for p1 <= mid && p2 <= hi {
		if nums[p1] <= nums[p2] {
			tmp[index] = nums[p1]
			p1++
		} else {
			tmp[index] = nums[p2]
			p2++
			// 统计逆序对的个数
			*cnt += mid - p1 + 1
		}
		index++
	}

	// 左边和右边有一个非空
	for p1 <= mid {
		tmp[index] = nums[p1]
		index++
		p1++
	}
	for p2 <= hi {
		tmp[index] = nums[p2]
		index++
		p2++
	}

	for i := 0; i < len(tmp); i++ {
		nums[i+low] = tmp[i]
	}
}
