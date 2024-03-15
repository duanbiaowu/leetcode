package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node
type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)

		row := make([]int, length)
		for i := 0; i < length; i++ {
			row[i] = queue[i].Val

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, row)

		queue = queue[length:]
	}

	return res
}

func levelOrderDFS(root *TreeNode) [][]int {
	var res [][]int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	if depth >= len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	dfs(root.Left, depth+1, res)
	dfs(root.Right, depth+1, res)
}
