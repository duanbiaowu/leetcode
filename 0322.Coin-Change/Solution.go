package leetcode

func coinChange(coins []int, amount int) int {
	// 需要硬币最多的情况就是全部由面额 1 的硬币组成
	// + 1 是为了初始化 dp 状态转移的上限值
	max := amount + 1

	// 初始化动态规划状态转移数组
	// 状态表达式: F(i)=min j=0…n−1 ​F(i−cj​)+1
	// dp[i] 表示金额为 i 时，所需要的最小硬币数量
	// 其中，cj 表示第 j 枚硬币的面值
	//      枚举最后一枚硬币面额是 cj ，那么需要从 i−cj ​ 这个金额的状态 F(i−cj) 转移过来
	//		再算上枚举的这枚硬币贡献
	// 由于要硬币数量最少，所以 F(i) 为前面能转移过来的状态的最小值，加上当前枚举的硬币数量 1
	dp := make([]int, max)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 只有硬币金额小于当前金额，才可以拼凑
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
