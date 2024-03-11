package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	for i := range inorder {
		if inorder[i] == root.Val {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
		}
	}

	return root
}

func buildTreeIteratively(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var stack []*TreeNode
	stack = append(stack, root)

	index := 0
	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		if node.Val != inorder[index] {
			node.Left = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index++
			}
			node.Right = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Right)
		}
	}

	return root
}
