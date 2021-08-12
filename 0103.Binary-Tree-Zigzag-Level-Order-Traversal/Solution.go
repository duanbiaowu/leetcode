package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for level := 0; len(queue) > 0; level++ {
		count := len(queue)
		res = append(res, make([]int, count))

		for i := 0; i < count; i++ {
			res[level][i] = queue[0].Val
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}

		// 偶数层级反转数组
		if level&1 == 1 {
			for left, right := 0, len(res[level])-1; left < right; {
				res[level][left], res[level][right] = res[level][right], res[level][left]
				left++
				right--
			}
		}
	}

	return res
}
