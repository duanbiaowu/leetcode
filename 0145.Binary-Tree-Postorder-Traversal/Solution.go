package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorderTraversalRecursively(root, &res)
	return res
}

func postorderTraversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		postorderTraversalRecursively(root.Left, res)
		postorderTraversalRecursively(root.Right, res)
		*res = append(*res, root.Val)
	}
}
