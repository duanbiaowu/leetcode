package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	// 当前节点值小于等于目标值，那么当前目标值的后继者必然在右子树
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	// 否则结果有可能是当前节点，或者在当前节点的左子树中
	// 那么先去左子树找一下试试，找不到的话当前节点即是结果
	node := inorderSuccessor(root.Left, p)
	if node == nil {
		return root
	}
	return node
}

func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	// 有右子树, 直接返回右子树最左边节点
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	// 没有右子树，向上回溯[第一个左转的祖先节点]
	// 比如查找节点 2，向上回溯到节点 1
	// 节点 1 是右转的，继续回溯到节点 3
	// 节点 3 是左转的，结束并返回
	//       3
	//      / \
	//     1   4
	//      \
	//       2
	var pre *TreeNode
	for root.Val != p.Val {
		if p.Val > root.Val {
			root = root.Right
		} else {
			pre = root
			root = root.Left
		}
	}
	return pre
}
