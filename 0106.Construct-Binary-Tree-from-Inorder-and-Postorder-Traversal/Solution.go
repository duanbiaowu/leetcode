package leetcode

import "github.com/duanbiaowu/leetcode/structures"

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
			root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
			root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):])
		}
	}
	return root
}
