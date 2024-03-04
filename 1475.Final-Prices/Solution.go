package leetcode

func finalPrices(prices []int) []int {
	n := len(prices)
	res := make([]int, n)
	// 栈顶始终保留一个值 0 作为 “哨兵”
	// 避免边界检查，语义表示 “没有折扣”
	stack := []int{0}

	for i := n - 1; i >= 0; i-- {
		// 如果栈顶元素比当前元素大，说明不符合折扣条件，直接弹出栈顶元素即可
		for len(stack) > 1 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		// 此时栈顶元素就是当前可以元素可以享受的折扣
		res[i] = prices[i] - stack[len(stack)-1]
		// 将当前元素入栈
		stack = append(stack, prices[i])
	}

	return res
}
