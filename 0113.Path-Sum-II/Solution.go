package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	dfs(root, &path, targetSum, &res)
	return res
}

func dfs(root *TreeNode, path *[]int, sum int, res *[][]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	// 从根节点到叶子节点, 路径总和等于给定目标和的路径
	// 如果总和相等, 但不是叶子节点，不算作路径
	if root.Val == sum && root.Left == nil && root.Right == nil {
		*res = append(*res, append([]int{}, *path...))
	}
	dfs(root.Left, path, sum-root.Val, res)
	dfs(root.Right, path, sum-root.Val, res)
	*path = (*path)[:len(*path)-1]
}
