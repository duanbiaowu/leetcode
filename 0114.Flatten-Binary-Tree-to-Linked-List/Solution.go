package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}

func flattenIteratively(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			prev := next
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}
