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

func isValidBST2(root *TreeNode) bool {
	inOrder := math.MinInt64

	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inOrder {
			return false
		}

		inOrder = root.Val
		root = root.Right
	}

	return true
}

// 二叉搜索树中序遍历满足有序性
// 通过中序遍历将值保存到数组中
// 然后判断是否有序即可
func isValidBSTInOrder(root *TreeNode) bool {
	var res []int
	inOrder(root, &res)

	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}

	return true
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}
