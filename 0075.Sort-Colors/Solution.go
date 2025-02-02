package leetcode

func sortColors(nums []int) {
	n := len(nums)
	p := 0

	// 首先分类编号为 0 的球
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}

	// 然后分类编号为 1 的球
	// 最后剩下的就是分类编号为 2 的球
	for i := p; i < n; i++ {
		if nums[i] == 1 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
}

func sortColors2(nums []int) {
	p1, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
