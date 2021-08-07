package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	traversalRecursively(root, &res)
	return res
}

func traversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		traversalRecursively(root.Left, res)
		traversalRecursively(root.Right, res)
	}
}
