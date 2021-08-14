package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func generateTrees(n int) []*TreeNode {
	return backtrack(1, n)
}

func backtrack(start, end int) []*TreeNode {
	// 两种极端情况: 左子树为nil  右子树为nil
	// 所以不能直接返回nil，影响回溯遍历
	if start > end {
		return []*TreeNode{nil}
	}

	var res []*TreeNode
	// 枚举可行根节点
	// 左子树的所有值小于根节点，右子树的所有值大于根节点
	for i := start; i <= end; i++ {
		leftTress := backtrack(start, i-1)
		rightTress := backtrack(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for left := range leftTress {
			for right := range rightTress {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  leftTress[left],
					Right: rightTress[right],
				})
			}
		}
	}

	return res
}
