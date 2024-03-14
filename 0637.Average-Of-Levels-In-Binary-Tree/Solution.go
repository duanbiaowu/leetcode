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
	var sum, cnt []int
	dfs(root, 0, &sum, &cnt)

	var res []float64
	for i := range sum {
		res = append(res, float64(sum[i])/float64(cnt[i]))
	}

	return res
}

func dfs(root *TreeNode, level int, sum, cnt *[]int) {
	if root == nil {
		return
	}

	// 遇到各个层级的第一个节点时，初始化累加器和计数器
	if level == len(*sum) {
		// 树的每一层，这个分支只会执行一次
		*sum = append(*sum, root.Val)
		*cnt = append(*cnt, 1)
	} else {
		// 树的每一层，这个分支会执行 N - 1 次
		// 其中 N 等于该层上面的节点数量
		(*sum)[level] += root.Val
		(*cnt)[level]++
	}

	dfs(root.Left, level+1, sum, cnt)
	dfs(root.Right, level+1, sum, cnt)
}
