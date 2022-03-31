package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
