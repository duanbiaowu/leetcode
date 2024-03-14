package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func kthSmallest2(root *TreeNode, k int) int {
	var res int

	dfs(root, &k, &res)

	return res
}

func dfs(root *TreeNode, k, res *int) {
	if root == nil {
		return
	}

	// 中序遍历模板
	// dfs (root.Left)
	// do something
	// dfs (root.Right)

	dfs(root.Left, k, res)

	*k--
	if *k == 0 {
		*res = root.Val
		return
	}

	dfs(root.Right, k, res)
}
