package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalRecursively(root, &res)
	return res
}

func inorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}

	return res
}

func inorderTraversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		inorderTraversalRecursively(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalRecursively(root.Right, res)
	}
}
