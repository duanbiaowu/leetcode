package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	dfs(root, &diameter)
	return diameter
}

func dfs(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	// 计算左右子树的深度
	left, right := 0, 0
	if root.Left != nil {
		left = dfs(root.Left, diameter) + 1
	}
	if root.Right != nil {
		right = dfs(root.Right, diameter) + 1
	}

	// 更新最大直径
	*diameter = max(*diameter, left+right)

	// 返回当前节点的深度 (以左右子树中深度较大的为准)
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
