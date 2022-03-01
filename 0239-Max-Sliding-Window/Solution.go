package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 || n-k < -1 {
		return []int{}
	}

	// 滑动窗口对应的数据结构为 双端队列
	var queue []int
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	// 未形成窗口（前 k 个数）
	for i := 0; i < k; i++ {
		push(i)
	}

	res := make([]int, len(nums)-k+1)
	res[0] = nums[queue[0]]
	// 形成窗口后
	for i := k; i < n; i++ {
		push(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		res[i-k+1] = nums[queue[0]]
	}
	return res
}
