package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	traversalRecursively(root, &res)
	return res
}

func preorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return res
}

func traversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		traversalRecursively(root.Left, res)
		traversalRecursively(root.Right, res)
	}
}
