package leetcode

	func maximumGap(nums []int) int {
		n := len(nums)
		if n < 2 {
			return 0
		}

		bucket := make([]int, n)
		maxVal := max(nums...)
		for exp := 1; exp <= maxVal; exp *= 10 {
			cnt := [10]int{}
			for i := 0; i < n; i++ {
				digit := nums[i] / exp % 10
				cnt[digit]++
			}
			for i := 1; i < 10; i++ {
				cnt[i] += cnt[i-1]
			}
			for i := n - 1; i >= 0; i-- {
				digit := nums[i] / exp % 10
				bucket[cnt[digit]-1] = nums[i]
				cnt[digit]--
			}
			copy(nums, bucket)
		}

		var res int
		for i := 1; i < n; i++ {
			res = max(res, nums[i]-nums[i-1])
		}
		return res
	}

	// 注意参数和遍历方式
	func max(nums ...int) int {
		res := nums[0]
		for _, v := range nums[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
