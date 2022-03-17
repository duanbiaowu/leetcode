package leetcode

// reference: https://leetcode-cn.com/problems/counting-bits/comments/44989
// i & (i - 1) 可以去掉 i 最右边的一个 1
// i & (i - 1）比 i 小
// 并且 i & (i - 1) 的 1 的个数已经在前面算过了
// 所以 i 的 1 的个数就是 i & (i - 1) 的 1 的个数 + 1
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

// reference: https://leetcode-cn.com/problems/counting-bits/comments/56708
// 一个二进制数，如果最低位为 1(%2 为 1）,则它与 n/2 的 1 个数相差1。
// 如果最低位为 0，则它与 n/2 的 1 个数相同
func countBits2(n int) []int {
	dp := make([]int, n+1)
	for i := 0; i <= n>>1; i++ {
		dp[i*2] = dp[i]
		if i*2+1 <= n {
			dp[i*2+1] = dp[i] + 1
		}
	}
	return dp
}
