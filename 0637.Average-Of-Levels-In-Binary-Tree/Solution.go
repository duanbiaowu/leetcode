package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// TreeNode Definition for a binary tree node
type TreeNode = structures.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sum := 0
		length := len(queue)

		for i := 0; i < length; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]

		res = append(res, float64(sum)/float64(length))
	}

	return res
}

func averageOfLevelsDFS(root *TreeNode) []float64 {
	// sum 存储每层节点累加值
	// cnt 存储每层节点数量
	var sum, cnt []int

	// 从 root 节点作为 DFS 起始节点
	// 使用指针传递状态变量
	dfs(root, 0, &sum, &cnt)

	// 计算每层的平均值
	var res []float64
	for i := range sum {
		res = append(res, float64(sum[i])/float64(cnt[i]))
	}

	return res
}

func dfs(root *TreeNode, depth int, sum, cnt *[]int) {
	// DFS 递归条件处理
	if root == nil {
		return
	}

	// 遇到每个层级的第一个节点时，初始化累加器和计数器
	if depth == len(*sum) {
		*sum = append(*sum, root.Val)
		*cnt = append(*cnt, 1)
	} else {
		// 遇到了每个层级的第 2 - N 个节点
		// 其中 N 等于该层上面的节点数量
		// 对于树的每一层，这个分支会执行 N - 1 次
		(*sum)[depth] += root.Val
		(*cnt)[depth]++
	}

	// 递归左子树时，深度 + 1, 同时传递两个指针数组
	dfs(root.Left, depth+1, sum, cnt)

	// 递归右子树时，深度 + 1, 同时传递两个指针数组
	dfs(root.Right, depth+1, sum, cnt)
}
