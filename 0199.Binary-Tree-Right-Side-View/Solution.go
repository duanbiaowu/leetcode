package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}

	return res
}
