package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

// 设交集链表长c,链表1除交集的长度为a，链表2除交集的长度为b:
// a + c + b = b + c + a
// 若无交集，则a + b = b + a, 返回 null
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 反向指向: 任一链表到尾节点时, 指向对方链表头节点
	// 消除链接长度差
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
