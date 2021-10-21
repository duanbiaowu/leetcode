package leetcode

// https://leetcode-cn.com/problems/jump-game/solution/55-by-ikaruga/
// 依次尝试每个起跳点，同时更新最远距离
func canJump(nums []int) bool {
	for i, k := 0, 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
