package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// ListNode Definition for singly-linked list.
type ListNode = structures.ListNode

// 设交集链表长 c, 链表 1 除交集的长度为 a，链表 2 除交集的长度为 b:
// a + c + b = b + c + a
// 若无交集，则 a + b = b + a, 返回 null
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 反向指向: 任一链表到尾节点时, 指向对方链表头节点
	// 消除链接长度差
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			// 注意这里需要指向链表 B 的头节点
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			// 注意这里需要指向链表 A 的头节点
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
