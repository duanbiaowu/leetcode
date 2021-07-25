package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	dump := &ListNode{0, head}
	fast, slow := head, dump
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dump.Next
}
