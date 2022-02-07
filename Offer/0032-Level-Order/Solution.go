package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		res = append(res, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
		queue = queue[1:]
	}

	return res
}
