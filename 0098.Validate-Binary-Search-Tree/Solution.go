package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"math"
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
