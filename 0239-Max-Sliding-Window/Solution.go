package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}

	// 滑动窗口对应的数据结构为 双端队列
	var deQueue []int
	push := func(i int) {
		for len(deQueue) > 0 && nums[i] >= nums[deQueue[len(deQueue)-1]] {
			deQueue = deQueue[:len(deQueue)-1]
		}
		deQueue = append(deQueue, i)
	}

	// 未形成窗口（前 k 个数）
	for i := 0; i < k; i++ {
		push(i)
	}

	res := make([]int, len(nums)-k+1)
	res[0] = nums[deQueue[0]]

	// 形成窗口后
	for i := k; i < len(nums); i++ {
		push(i)

		// 保证当前窗口大小和索引对应关系
		for deQueue[0] <= i-k {
			deQueue = deQueue[1:]
		}

		res[i-k+1] = nums[deQueue[0]]
	}

	return res
}

// 暴力法: 超出时间限制
func maxSlidingWindowSimple(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	var res []int
	for i := 0; i <= len(nums)-k; i++ {
		maxCur := nums[i]

		for j := i; j < i+k; j++ {
			if nums[j] > maxCur {
				maxCur = nums[j]
			}
		}

		res = append(res, maxCur)
	}

	return res
}
