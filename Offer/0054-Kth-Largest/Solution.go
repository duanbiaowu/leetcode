package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func kthLargest(root *TreeNode, k int) int {
	n, res := k, -1
	inorder(root, &n, &res)
	return res
}

func inorder(root *TreeNode, k, res *int) {
	if root == nil {
		return
	}
	inorder(root.Right, k, res)
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}
	inorder(root.Left, k, res)
}

func kthLargest2(root *TreeNode, k int) int {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return cur.Val
			}
			cur = cur.Left
		}
	}
	return -1
}
