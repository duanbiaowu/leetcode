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

func getMinimumDifference2(root *TreeNode) int {
	vals := inorderTraversal(root)
	if len(vals) == 0 {
		return -1
	}

	minDiff := math.MaxInt64
	for i := 1; i < len(vals); i++ {
		minDiff = min(minDiff, vals[i]-vals[i-1])
	}

	return minDiff
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
