package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return checkSubTree(t1.Left, t2.Left) && checkSubTree(t2.Right, t2.Right)
	}
	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}
