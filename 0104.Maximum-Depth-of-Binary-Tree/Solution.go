package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func dfs(root *TreeNode, depth int, res *int) {
	if root != nil {
		*res = max(*res, depth+1)
		dfs(root.Left, depth+1, res)
		dfs(root.Right, depth+1, res)
	}
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	queue = append(queue, root)
	depth := 0

	for n := len(queue); n > 0; n = len(queue) {
		depth++
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
