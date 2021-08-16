package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func isSymmetricBFS(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left)
		queue = append(queue, q.Right)
		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}

	return true
}
