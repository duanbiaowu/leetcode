package leetcode

import "leetcode/structures"

// ListNode structures.ListNode
type ListNode = structures.ListNode

func getLength(head *ListNode) int {
	length := 0
	for head.Next != nil {
		head = head.Next
		length++
	}
	return length
}

// simple solution
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if n <= 0 {
//		return head
//	}
//
//	dump := &ListNode{0, head}
//	length := getLength(head)
//	cur := dump
//	for i := 0; i < length-n+1; i++ {
//		cur = cur.Next
//	}
//	cur.Next = cur.Next.Next
//	return dump.Next
//}

// fast-slow pointer solution
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	// 哨兵节点
	dummy := &ListNode{Val: 0, Next: head}
	// 注意快慢指针的初始化指向
	fast, slow := head, dummy

	// 快指针先走 N 步
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	// 然后快慢指针一起走
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 当快指针扫描到链表结尾时
	// 慢指针指向的正好是需要删除的节点的前置节点
	slow.Next = slow.Next.Next

	return dummy.Next
}

// recursive solution
func removeNthFromEndRecursively(head *ListNode, n int) *ListNode {
	counter := 0
	return removeNthFromEndRecursivelyInternal(head, n, &counter)
}

func removeNthFromEndRecursivelyInternal(head *ListNode, n int, counter *int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}

	head.Next = removeNthFromEndRecursivelyInternal(head.Next, n, counter)
	*counter++
	if n == *counter {
		return head.Next
	}
	return head
}
