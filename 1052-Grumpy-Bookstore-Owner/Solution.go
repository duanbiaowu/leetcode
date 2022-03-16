package leetcode

// reference: https://leetcode-cn.com/problems/grumpy-bookstore-owner/solution/yong-mi-mi-ji-qiao-wan-liu-zhu-zui-duo-d-py41/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	sum, n := 0, len(customers)
	if n == 0 || n != len(grumpy) {
		return 0
	}

	// 老板不生气时顾客数
	for i := range customers {
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}

	// 生气的 minutes 分钟内，会让多少顾客不满意
	cnt := 0
	// 先计算起始的 [0, minutes) 区间
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			cnt += customers[i]
		}
	}
	res := cnt

	// 然后利用滑动窗口，每次向右移动一步
	for i := minutes; i < n; i++ {
		// 如果新进入窗口的元素是生气的，累加不满意的顾客到滑动窗口中
		if grumpy[i] == 1 {
			cnt += customers[i]
		}
		// 如果离开窗口的元素是生气的，则从滑动窗口中减去该不满意的顾客数
		if grumpy[i-minutes] == 1 {
			cnt -= customers[i-minutes]
		}
		// 求所有窗口内不满意顾客的最大值
		res = max(res, cnt)
	}

	// 老板不生气时顾客数 + 窗口 minutes 内挽留的因为生气被赶走的顾客数
	return sum + res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
