package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

// 中序遍历
func convertBiNode(root *TreeNode) *TreeNode {
	// 哨兵节点
	head := &TreeNode{}
	pre := head
	cur := root
	var stack []*TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 左指针置空
			cur.Left = nil
			// 指针作为链表的next指针
			pre.Right = cur
			// 指针后移
			pre = cur
			cur = cur.Right
		}
	}
	return head.Right
}

func convertBiNodeRecursively(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	inorder(root, head)
	return head.Right
}

// 中序递归遍历
func inorder(root, pre *TreeNode) *TreeNode {
	if root != nil {
		pre = inorder(root.Left, pre)
		root.Left = nil
		pre.Right = root
		pre = root
		pre = inorder(root.Right, pre)
	}

	return pre
}
