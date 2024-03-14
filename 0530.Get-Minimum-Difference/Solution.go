package leetcode

import (
	"math"

	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func getMinimumDifference(root *TreeNode) int {
	res, prev := math.MaxInt64, -1

	dfs(root, &res, &prev)

	return res
}

func dfs(root *TreeNode, res, prev *int) {
	if root == nil {
		return
	}

	// 中序遍历模板
	// dfs (root.Left)
	// do something
	// dfs (root.Right)

	dfs(root.Left, res, prev)

	if *prev != -1 {
		*res = min(*res, root.Val-*prev)
	}
	*prev = root.Val

	dfs(root.Right, res, prev)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
