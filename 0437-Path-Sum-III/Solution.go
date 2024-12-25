package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	cnt := rootSum(root, targetSum)
	cnt += pathSum(root.Left, targetSum)
	cnt += pathSum(root.Right, targetSum)
	return cnt
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	cnt := 0
	if root.Val == targetSum {
		cnt++
	}
	cnt += rootSum(root.Left, targetSum-root.Val)
	cnt += rootSum(root.Right, targetSum-root.Val)
	return cnt
}

// 前缀和
func pathSum2(root *TreeNode, targetSum int) int {
	preSum := make(map[int]int)
	// 重点: 初始化前缀和为 0 的路径为 1
	// 针对某个节点的前缀和为 sum
	// 如果存在一个前缀和为 sum - targetSum 的路径，那么这个路径的和就是 targetSum
	preSum[0] = 1

	cnt := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		sum += root.Val
		cnt += preSum[sum-targetSum]
		preSum[sum]++

		dfs(root.Left, sum)
		dfs(root.Right, sum)

		preSum[sum]--
	}

	dfs(root, 0)

	return cnt
}
