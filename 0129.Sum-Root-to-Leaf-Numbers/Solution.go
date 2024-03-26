package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func sumNumbers(root *TreeNode) int {
	// 从 root 节点作为 DFS 起始节点
	// “当前累加值” 初始化为 0
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	// DFS 递归条件处理
	if root == nil {
		return 0
	}

	// 计算并更新 “当前累加值”
	sum = sum*10 + root.Val

	// 如果当前节点的左右子节点都为 nil, 说明当前节点就是叶子节点
	// 直接返回 “当前累加值”
	if root.Left == nil && root.Right == nil {
		return sum
	}

	// 递归调用左右子树 (并传入 “当前累加值” )
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func bfs(root *TreeNode) int {
	// 边界处理
	if root == nil {
		return 0
	}

	// 返回值, 路径总和
	sum := 0

	// 初始化两个队列
	// 将根节点加入到节点队列中
	nodeQueue := []*TreeNode{root}
	// 将根节点的值加入到 路径数字总和队列 中
	numQueue := []int{root.Val}

	for len(nodeQueue) > 0 {
		// 获取当前节点
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		// 获取当前 “路径和”
		num := numQueue[0]
		numQueue = numQueue[1:]

		if node.Left == nil && node.Right == nil {
			// 叶子节点，直接将当前的 “路径和” 累加到返回值总和中
			sum += num
		} else {
			// 将当前节点的左右子节点加入队列
			// 将 当前 “路径和” 和当前节点的左右子节点的值 重新计算
			//   然后加入 路径数字总和队列
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
