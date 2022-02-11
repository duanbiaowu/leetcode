package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func isBalanced(root *TreeNode) bool {
	return heightB2T(root) >= 0
	//if root == nil {
	//	return true
	//}
	//return abs(heightT2B(root.Left) - heightT2B(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 自顶向下递归: 对于同一个节点, 会重复计算
// 优化点: 改为自底向上
func heightT2B(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(heightT2B(root.Left), heightT2B(root.Right)) + 1
}

// 自底向上递归:
// 1. 对于当前节点, 先递归判断其左右子树是否平衡
// 2. 再判断以当前节点为根的子树是否平衡
// 3. 如果一颗子树是平衡的，返回其高度(高度一定是非负数), 否则返回-1
func heightB2T(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := heightB2T(root.Left)
	rightHeight := heightB2T(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
