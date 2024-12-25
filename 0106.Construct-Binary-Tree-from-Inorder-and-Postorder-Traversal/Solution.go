package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// TreeNode Definition for a binary tree node.
type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	// 后序遍历数组的最后一个元素就是根节点
	root := &TreeNode{Val: postorder[n-1]}
	// 删除最后一个元素 (因为已使用)
	postorder = postorder[:n-1]

	for i := range inorder {
		// 找到中序遍历数组中根节点的值对应的索引
		// 并根据索引将数组分成两个部分
		if inorder[i] == root.Val {
			// 左半部分就是左子树的所有节点值
			root.Left = buildTree(inorder[:i], postorder[:i])
			// 右半部分就是右子树的所有节点值
			// 因为中序遍历结果中的当前元素已经作为本轮递归返回的根节点
			// 所以这里 inorder 的索引需要加 1 作为右子树的参数
			root.Right = buildTree(inorder[i+1:], postorder[i:])
		}
	}

	return root
}
