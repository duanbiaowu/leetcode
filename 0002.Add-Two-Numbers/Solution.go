package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	num1, num2, carry := 0, 0, 0
	cur := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		// 计算链表 1 的当前值
		if l1 == nil {
			num1 = 0
		} else {
			num1 = l1.Val
			l1 = l1.Next
		}

		// 计算链表 2 的当前值
		if l2 == nil {
			num2 = 0
		} else {
			num2 = l2.Val
			l2 = l2.Next
		}

		// 当计算出来的新节点放到当前节点后面
		cur.Next = &ListNode{
			Val: (num1 + num2 + carry) % 10,
		}
		// 计算是否需要进位
		carry = (num1 + num2 + carry) / 10
		// 更新当前节点
		cur = cur.Next
	}

	return dummy.Next
}
