package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	cnt := rootSum(root, targetSum)
	cnt += pathSum(root.Left, targetSum)
	cnt += pathSum(root.Right, targetSum)
	return cnt
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	cnt := 0
	if root.Val == targetSum {
		cnt++
	}
	cnt += rootSum(root.Left, targetSum-root.Val)
	cnt += rootSum(root.Right, targetSum-root.Val)
	return cnt
}

// 前缀和
func pathSum2(root *TreeNode, targetSum int) int {
	res := 0
	preSum := make(map[int]int)
	preSum[0] = 1

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}
		cnt += root.Val
		res += preSum[cnt-targetSum]
		preSum[cnt]++
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
		preSum[cnt]--
	}

	dfs(root, 0)
	return res
}
