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

func maxPathSum2(root *TreeNode) int {
	maxSum := math.MinInt32
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// 子节点路径和为负数时，剪枝
	left := max(0, dfs(root.Left, maxSum))
	right := max(0, dfs(root.Right, maxSum))

	// 更新最大路径和对应的全局变量
	*maxSum = max(*maxSum, root.Val+left+right)

	// 最大路径只能是根节点 + 左右节点中的任意一个
	// 所以必须从左右节点中选择一个节点出来
	return root.Val + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
