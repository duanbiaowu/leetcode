package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalRecursively(root, &res)
	return res
}

func inorderTraversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		inorderTraversalRecursively(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalRecursively(root.Right, res)
	}
}
