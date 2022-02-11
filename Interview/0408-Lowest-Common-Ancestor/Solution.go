package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

// 非二叉搜索树
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// 递归出口：根节点为空，或者找到 p 或 q
	if root == nil || p == root || q == root {
		return root
	}
	// 递归式：从根节点的左右子树寻找p、q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 判断祖先位置
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}
