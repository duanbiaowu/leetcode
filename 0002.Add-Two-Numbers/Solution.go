package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		// 计算链表 1 的当前值
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		// 计算链表 2 的当前值
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 当计算出来的新节点放到当前节点后面
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		// 计算是否需要进位
		carry = (n1 + n2 + carry) / 10
		// 更新当前节点
		current = current.Next
	}

	return head.Next
}
