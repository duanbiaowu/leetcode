package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

type TreeNode = structures.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	return build(head, nil)
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
