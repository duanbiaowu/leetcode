package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var dfs func(A *TreeNode, B *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil {
			return false
		}
		return A.Val == B.Val && dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
	}

	if A == nil || B == nil {
		return false
	}
	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
