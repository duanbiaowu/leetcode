package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	// 题目说明：
	// 	  1. 每个节点的子节点数量只能为 2 或 0
	//    2. 如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个
	// 所以最小的值一定是根结点的值
	// 找到第一个比根节点大的值立即返回，不用再递归下去
	// 最终问题转化为求左右子树的最小值
	return findFirstBigger(root, root.Val)
}

func findFirstBigger(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val > val {
		return root.Val
	}

	left := findFirstBigger(root.Left, val)
	right := findFirstBigger(root.Right, val)

	// 如果左右子树最小值都大于根节点的值，取两者较小值
	if left > 0 && right > 0 {
		return min(left, right)
	}
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
