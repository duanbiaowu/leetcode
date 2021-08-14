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
	for i := start; i <= end; i++ {
		leftTress := backtrack(start, i-1)
		rightTress := backtrack(i+1, end)
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
