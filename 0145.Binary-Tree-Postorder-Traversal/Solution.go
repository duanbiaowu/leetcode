package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorderTraversalRecursively(root, &res)
	return res
}

func postorderTraversalRecursively(root *TreeNode, res *[]int) {
	if root != nil {
		postorderTraversalRecursively(root.Left, res)
		postorderTraversalRecursively(root.Right, res)
		*res = append(*res, root.Val)
	}
}

func postorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var prev *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果不存在右孩子节点 或者 右孩子节点 已经被访问过
		// 直接访问当前节点
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			node = nil
		} else {
			// 当前节点再次入栈 优先访问其右孩子节点
			stack = append(stack, node)
			node = node.Right
		}
	}

	return res
}

func postorderTraversalIteratively2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	var prev, cur *TreeNode
	stack = append(stack, root)

	// 要保证根结点在左孩子和右孩子访问之后才能访问，因此对于任一结点P，先将其入栈
	// 如果P不存在左孩子和右孩子，则可以直接访问该结点
	// 或者如果P存在左孩子或者右孩子，但是其左孩子和右孩子都已被访问过了，则同样可以直接访问该结点

	// 若非上述两种情况，则将P的右孩子和左孩子依次入栈，这样就保证了每次取栈顶元素的时候
	// 左孩子在右孩子前面被访问，左孩子和右孩子都在根结点前面被访问
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) ||
			(prev != nil && (prev == cur.Left || prev == cur.Right)) {
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}

	return res
}
