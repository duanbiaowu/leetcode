package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

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
	res := 0

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
