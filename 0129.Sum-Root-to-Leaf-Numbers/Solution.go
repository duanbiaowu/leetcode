package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	nodeQueue := []*TreeNode{root}
	numQueue := []int{root.Val}

	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		num := numQueue[0]
		numQueue = numQueue[1:]

		if node.Left == nil && node.Right == nil {
			sum += num
		} else {
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
				numQueue = append(numQueue, num*10+node.Left.Val)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
				numQueue = append(numQueue, num*10+node.Right.Val)
			}
		}
	}

	return sum
}
