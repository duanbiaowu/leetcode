package leetcode

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	// 从倒数第二天开始，倒数第一天必然是 0
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		// 两个变量模拟单调栈
		for j < n {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			} else if res[j] == 0 {
				break
			} else {
				j += res[j]
			}
		}
	}
	return res
}

func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return res
}
