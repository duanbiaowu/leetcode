package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

// ListNode Definition for singly-linked list.
type ListNode = structures.ListNode

// 求交集问题，首先要找到长度差
// 从终点反过来看问题，这个问题的本质就会变得很简单：
// 如果两个链表有交集，那么不管两个链表的长度是多少，其最终指向的末尾节点 (倒数第二个节点) 都是一样的

// 所以针对两个链表长度不同的情况，只需要在一方链表指针到达末尾时，将指针指向另一方链表的头节点，
// 这样经过两轮遍历: 两个链表的长度就一样的:
//   如果存在交集节点，那么在第二轮必然会 “相遇”，然后返回即可
//   如果不存在交集节点，链表指针会一直扫描到末尾 (返回 nil)

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
