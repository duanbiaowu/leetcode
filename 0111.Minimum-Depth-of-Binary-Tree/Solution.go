package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left > 0 && right > 0 {
		return min(left, right) + 1
	}
	// 其中一个子树深度为0, 无需知道具体是哪个，相加即可, 不影响结果
	return left + right + 1
}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)

	// 队列Top元素不出队的实现
	// 结果一致, 但降低了可读性
	level := 1
	for i := 0; i < len(queue); {
		n := len(queue)

		for ; i < n; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}
	return level
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
