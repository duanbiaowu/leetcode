package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxLen := 0
	// 状态转移方程: dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
	// 其中 nums[j] < nums[i]
	// dp[i] 的值代表以 nums[i] 结尾的最长子序列长度
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		// 默认为 1
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

// reference: https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-2/
func lengthOfLIS2(nums []int) int {
	// tails[k] 的值代表 长度为 k+1 子序列 的尾部元素值
	tails := make([]int, len(nums))
	res := 0

	for i := range nums {
		low, hi := 0, res
		for low < hi {
			mid := low + (hi-low)>>1
			if tails[mid] < nums[i] {
				low = mid + 1
			} else {
				hi = mid
			}
		}
		tails[low] = nums[i]
		if hi == res {
			res++
		}
	}

	return res
}

// 从数组中的每个元素开始，尝试找出以该元素开头的所有递增子序列
// 对于每个元素，可以选择将其包含在递增子序列中或者不包含，因此可以通过递归实现
// 在递归的过程中，记录当前递增子序列的长度，并更新最长递增子序列的长度

// 暴力搜索迭代版本 (超时)
// 超时原因: 重复检测
func lengthOfLISBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLen := 1
	backtrack(nums, -1, 0, &maxLen)

	return maxLen
}

// 回溯穷举每个元素开头的所有递增子序列
// prev: 上个元素的索引
// cur:  当前元素的索引
func backtrack(nums []int, prev, cur int, maxLen *int) int {
	if cur == len(nums) {
		return 0
	}

	taken := 0
	if prev == -1 || nums[cur] > nums[prev] {
		// 计算将当前元素包含在子序列中时
		// 可以获得的最大递增子序列长度
		taken = 1 + backtrack(nums, cur, cur+1, maxLen)
	}
	// 计算跳过当前元素时 (也就是不将当前元素包含在子序列中)
	// 可以获得的最大递增子序列长度
	notTaken := backtrack(nums, prev, cur+1, maxLen)

	// 更新已知的最大递增子序列长度
	*maxLen = max(*maxLen, max(taken, notTaken))

	return max(taken, notTaken)
}

func lengthOfLISMemo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化备忘录
	// 所有值设置为 -1
	// 表示当前元素组合还未被检测
	// 使用一个二维数组来表示 HashMap
	// 其中 memo[i][j] 表示:
	//    i = 参数 prev + 1 (prev 从 -1 开始)
	//    j = 参数 cur
	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, len(nums))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return backtrackMemo(nums, -1, 0, memo)
}

func backtrackMemo(nums []int, prev, cur int, memo [][]int) int {
	// 递归结束条件
	if cur == len(nums) {
		return 0
	}

	// 如果备忘录中已经包含 [当前元素组合]
	// 直接返回
	if memo[prev+1][cur] != -1 {
		return memo[prev+1][cur]
	}

	taken := 0
	if prev == -1 || nums[cur] > nums[prev] {
		// 计算将当前元素包含在子序列中时
		// 可以获得的最大递增子序列长度
		taken = 1 + backtrackMemo(nums, cur, cur+1, memo)
	}
	// 计算跳过当前元素时 (也就是不将当前元素包含在子序列中)
	// 可以获得的最大递增子序列长度
	notTaken := backtrackMemo(nums, prev, cur+1, memo)

	// 更新备忘录中 [当前元素组合] 的值
	memo[prev+1][cur] = max(taken, notTaken)

	return memo[prev+1][cur]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
