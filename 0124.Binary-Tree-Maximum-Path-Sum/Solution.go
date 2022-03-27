package leetcode

import (
	"math"

	"github.com/duanbiaowu/leetcode/structures"
)

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := math.MinInt32
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 子节点路径和为负数时，剪枝
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))

		// 更新最大路径和
		maxSum = max(maxSum, root.Val+left+right)

		return root.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
