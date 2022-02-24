package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}
	var stack []int
	j := 0
	for i := range pushed {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
