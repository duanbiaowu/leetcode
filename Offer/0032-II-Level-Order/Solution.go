package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for n := len(queue); n > 0; n = len(queue) {
		row := make([]int, n)
		for i := 0; i < n; i++ {
			row[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		res = append(res, row)
	}

	return res
}
