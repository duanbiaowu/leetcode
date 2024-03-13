package leetcode

import (
	"math"

	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, hi int) bool {
	if root == nil {
		return true
	}
	// 题目要求:
	// 节点的左子树只包含小于当前节点的数
	// 节点的右子树只包含大于当前节点的数
	if root.Val <= low || root.Val >= hi {
		return false
	}
	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, hi)
}

func isValidBSTInOrder(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
