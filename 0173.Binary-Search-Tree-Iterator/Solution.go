package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{[]*TreeNode{}}
	iterator.LeftMostInOrder(root)
	return iterator
}

func (this *BSTIterator) Next() int {
	top := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	if top.Right != nil {
		this.LeftMostInOrder(top.Right)
	}
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

// LeftMostInOrder
// 二叉搜索树的中序遍历序列是单调递增的
// 利用二叉树的迭代方式的中序遍历，保存左子链，从而使用O(h)的内存
func (this *BSTIterator) LeftMostInOrder(root *TreeNode) {
	for root != nil {
		this.Stack = append(this.Stack, root)
		root = root.Left
	}
}
