package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 反转后半部分
	var last *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = last
		last = slow
		slow = next
	}
	// 前半部分和后半部分对比
	first := head
	for last != nil {
		if first.Val != last.Val {
			return false
		}
		first = first.Next
		last = last.Next
	}
	return true
}
