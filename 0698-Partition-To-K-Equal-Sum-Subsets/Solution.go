package leetcode

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) == 0 || k == 0 {
		return false
	}
	if k == 1 {
		return true
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 剪枝
	if sum%k > 0 {
		return false
	}
	sort.Ints(nums)
	avg := sum / k
	// 剪枝
	if nums[len(nums)-1] > avg {
		return false
	}

	n := len(nums)

	// index: 搜索起始索引，因为不可能每次都去遍历所有的数组元素
	// cur:  当前子集合元素总和
	// cnt:  记录已经有多少个总和为 k 的组合数量
	// memo: 记录数组中哪些元素已经被使用
	var dfs func(index, cur, cnt int, memo []bool) bool
	dfs = func(index, cur, cnt int, memo []bool) bool {
		// 如果组合数量已经满足条件，剪枝
		if cnt == k {
			return true
		}
		// 如果当前总和满足等于平均数
		// 组合数量 + 1, 然后继续向前贪心选择
		if cur == avg {
			return dfs(n-1, 0, cnt+1, memo)
		}

		// 贪心策略: 倒序顺序
		// 每个回合的搜索总是在搜索「剩余未使用元素的最大值」所在的那个集合
		// 并且按照「优先使用大数值」的原则来构造
		for i := index; i >= 0; i-- {
			if memo[i] || cur+nums[i] > avg {
				// 可行性剪枝
				continue
			}

			// 标记当前元素已被使用
			memo[i] = true
			// 继续向前贪心选择
			if dfs(i-1, cur+nums[i], cnt, memo) {
				return true
			}
			// 标记当前元素未被使用
			memo[i] = false

			if cur == 0 {
				// 可行性剪枝
				// 如果当前元素不存在于任何满足条件的子集合中
				// 那么必然不满足题目要求的正好分成 k 个子集条件，直接返回即可
				return false
			}
		}

		return false
	}

	return dfs(n-1, 0, 0, make([]bool, n))
}
