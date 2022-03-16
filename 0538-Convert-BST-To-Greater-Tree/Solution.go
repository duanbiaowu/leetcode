package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return helper(root, &sum)
}

func helper(root *TreeNode, sum *int) *TreeNode {
	if root != nil {
		helper(root.Right, sum)
		*sum += root.Val
		root.Val = *sum
		helper(root.Left, sum)
	}

	return root
}
