package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func longestUniValuePath(root *TreeNode) int {
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

	// 如果当前节点和孩子节点不同，路径重置
	if left > 0 && root.Val != root.Left.Val {
		left = 0
	}
	if right > 0 && root.Val != root.Right.Val {
		right = 0
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
