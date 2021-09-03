package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

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
