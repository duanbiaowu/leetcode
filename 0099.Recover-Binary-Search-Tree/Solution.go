package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"math"
)

type TreeNode = structures.TreeNode

func recoverTree(root *TreeNode) {
	var (
		first   *TreeNode
		second *TreeNode
		prev   = &TreeNode{Val: math.MinInt32}
	)

	// 参数使用指向指针的指针
	// 仅使用指针递归时值丢失，结果错误
	inorder(root, &first, &second, &prev)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func inorder(root *TreeNode, first, second, prev **TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, first, second, prev)
	if prev != nil && (*prev).Val > root.Val {
		if *first == nil {
			*first = *prev
		}
		*second = root
	}

	*prev = root
	inorder(root.Right, first, second, prev)
}
