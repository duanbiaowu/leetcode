package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

// 解题的关键在于找到两个交换节点，然后交换过来
// 二叉搜索树的中序遍历是递增的
// 第一个节点: 遍历时前一个节点大于后一个节点（第一个）
// 第二个节点: 遍历时前一个节点大于后一个节点（第二个）
// PS: 	一般情况下，树的遍历递归版本比迭代版本更直观和易于理解
//		但由于递归版本使用了双重指针，所以要仔细看清楚指针在不同情况下指向的具体值
func recoverTree(root *TreeNode) {
	var (
		first  *TreeNode
		second *TreeNode
		prev   *TreeNode
	)

	// 参数使用指向指针的指针
	// 仅使用指针递归时值丢失，结果错误
	inorder(root, &first, &second, &prev)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func inorder(root *TreeNode, first, second, prev **TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, first, second, prev)
	if *prev != nil && (*prev).Val > root.Val {
		if *first == nil {
			*first = *prev
		}
		*second = root
	}

	*prev = root
	inorder(root.Right, first, second, prev)
}

func recoverTreeIteratively(root *TreeNode) {
	var (
		first  *TreeNode
		second *TreeNode
		prev   *TreeNode
	)

	var stack []*TreeNode
	for node := root; node != nil || len(stack) > 0; {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev
			}
			second = root
		}

		prev = node
		node = node.Right
	}

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
