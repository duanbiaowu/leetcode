package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

type TreeNode = structures.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	//return build(head, nil)

	// 注意: 这里使用双重指针，而非全局变量
	return buildTree(0, getLength(head)-1, &head)
}

func build(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = build(left, mid)
	root.Right = build(mid.Next, right)
	return root
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right int, head **ListNode) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left+1)>>1
	root := &TreeNode{}
	root.Left = buildTree(left, mid-1, head)
	root.Val = (*head).Val
	*head = (*head).Next
	root.Right = buildTree(mid+1, right, head)
	return root
}

func getLength(head *ListNode) int {
	var length int
	for length = 0; head != nil; head = head.Next {
		length++
	}
	return length
}
