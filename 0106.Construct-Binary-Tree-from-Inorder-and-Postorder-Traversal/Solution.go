package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[n-1]}
	postorder = postorder[:n-1]
	for i, val := range inorder {
		if val == root.Val {
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:])
		}
	}
	return root
}
