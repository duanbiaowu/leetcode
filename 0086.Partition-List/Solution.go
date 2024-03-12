package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	// 哨兵节点：指向 small 链表的头节点
	smallDummy := &ListNode{}
	// 哨兵节点：指向 large 链表的头节点
	largeDummy := &ListNode{}

	// 小于 X 的所有节点形成的链表
	small := smallDummy
	// 大于等于 X 的所有节点形成的链表
	large := largeDummy

	// 直接使用 head 作为游标指针
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	// 切断 large 链表的尾指针
	// 例如源链表为 1 -> 4 -> 3 -> 2 -> 5 -> 2, X = 3
	// 那么在分割完成后
	// small 链表为 1 -> 2 -> 2 -> 5 ...
	// large 链表为 4 -> 3 -> 5 -> 2 ...
	// 这时就需要将 large 链表末尾指针给切断，否则就形成了 “环形链表”
	large.Next = nil
	// 将 small 链表连接到 large 链表
	small.Next = largeDummy.Next

	// 返回 small 链表的头节点
	return smallDummy.Next
}
