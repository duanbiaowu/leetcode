package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func bfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		top.Left, top.Right = top.Right, top.Left
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}

	return root
}
