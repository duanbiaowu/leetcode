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

// fast-slow pointer solution
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
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
