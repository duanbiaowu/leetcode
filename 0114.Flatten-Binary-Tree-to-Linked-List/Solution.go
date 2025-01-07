package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归先将左右子树展开完成
	flatten(root.Right)
	flatten(root.Left)

	// 开始展开当前节点的左右子树
	right := root.Right
	root.Right = root.Left
	root.Left = nil

	// 将左子树中最右侧的节点 连接到 当前节点的右子树
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
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
