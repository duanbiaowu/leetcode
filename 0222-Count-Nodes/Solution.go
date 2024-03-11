package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

// reference: https://leetcode-cn.com/problems/count-complete-tree-nodes/solution/chang-gui-jie-fa-he-ji-bai-100de-javajie-fa-by-xia/
// 完全二叉树的定义：
//	一棵空树或者它的叶子节点只出在最后两层，
// 	若最后一层不满则叶子节点只在最左侧

// 满二叉树的层数为h，则总节点数为：2^h - 1
// 对 root 节点的左右子树进行高度统计，分别记为 left 和 right，有以下两种结果:

// left == right
//
//	说明左子树一定是满二叉树，因为节点已经填充到右子树了，左子树必定已经填满了。
//	所以左子树的节点总数我们可以直接得到，是 2^left - 1，
//	加上当前 root 节点，则正好是 2^left，
//	再对右子树进行递归统计。
//
// left != right
//
//	说明此时最后一层不满，但倒数第二层已经满了，
//	可以直接得到右子树的节点个数，
//	同理右子树节点 + root 节点，总数为 2^right，
//	再对左子树进行递归查找。
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := 0, 0
	node := root.Left

	// 计算左子树的高度
	// 以最左边的节点为准
	for node != nil {
		left++
		node = node.Left
	}

	node = root.Right
	// 计算右子树的高度
	// 以最左边的节点为准
	for node != nil {
		right++
		node = node.Left
	}

	if left == right {
		return countNodes(root.Right) + 1<<left
	}
	return countNodes(root.Left) + 1<<right
}

// 二叉树通用解法
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
