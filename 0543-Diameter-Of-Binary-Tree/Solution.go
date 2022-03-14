package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	var cnt int
	dfs(root, &cnt)
	return cnt
}

func dfs(root *TreeNode, cnt *int) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	}
	left, right := 0, 0
	if root.Left != nil {
		left = dfs(root.Left, cnt) + 1
	}
	if root.Right != nil {
		right = dfs(root.Right, cnt) + 1
	}
	*cnt = max(*cnt, left+right)
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
