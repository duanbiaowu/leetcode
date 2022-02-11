package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	sum -= root.Val
	if sum == 0 {
		count++
	}
	count += helper(root.Left, sum)
	count += helper(root.Right, sum)
	return count
}
