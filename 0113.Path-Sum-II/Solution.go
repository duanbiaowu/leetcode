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

func pathSumBFS(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	type pair struct {
		node *TreeNode
		sum  int
	}

	var res [][]int
	var parent = map[*TreeNode]*TreeNode{}
	var queue []pair
	queue = append(queue, pair{root, 0})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		node := top.node
		sum := top.sum + node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				res = append(res, getPath(parent, node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, sum})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, sum})
			}
		}
	}

	return res
}

func getPath(parent map[*TreeNode]*TreeNode, node *TreeNode) []int {
	var path []int
	for ; node != nil; node = parent[node] {
		path = append(path, node.Val)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
