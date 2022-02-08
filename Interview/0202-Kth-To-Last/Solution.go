package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	fast := head
	slow := head
	for i := 1; i < k && fast.Next != nil; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
