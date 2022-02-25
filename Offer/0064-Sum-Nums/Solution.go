package leetcode

// reference: https://leetcode-cn.com/problems/qiu-12n-lcof/solution/golangdi-gui-shuang-bai-by-daniel-53/
func sumNums(n int) int {
	_ = n > 0 && func() bool {
		n += sumNums(n - 1)
		return true
	}()
	return n
}
