package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// ListNode Definition for singly-linked list.
type ListNode = structures.ListNode

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p1, p2 := head1, head2
	cur := dummy

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			cur.Next = p2
			p2 = p2.Next
		} else {
			cur.Next = p1
			p1 = p1.Next
		}
		cur = cur.Next
	}

	if p1 != nil {
		cur.Next = p1
	} else {
		cur.Next = p2
	}

	return dummy.Next
}

func sortListIteratively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	for step := 1; step < length; step <<= 1 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			// 需要合并的链表块左边起始位置
			head1 := cur
			for i := 1; i < step && cur.Next != nil; i++ {
				cur = cur.Next
			}

			// 需要合并的链表块右边起始位置
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < step && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			// 需要合并的下一个链表块起始位置
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)
			// 合并后链表指针前进
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}

	return dummy.Next
}

// https://leetcode-cn.com/problems/sort-list/solution/sort-list-gui-bing-pai-xu-lian-biao-by-jyd/1082575
func quickSort(head, tail *ListNode) {
	// 注意: head 为 dummyNode
	if head == tail || head.Next == tail || head.Next.Next == tail {
		return
	}

	// 第一个节点为基准
	pivot := head.Next
	// 建立临时链表
	dummy := &ListNode{}
	left, right := dummy, pivot

	for right.Next != tail {
		// 如果当前元素小于基准，就加入临时链表，并在原链表中删除
		if right.Next.Val < pivot.Val {
			left.Next = right.Next
			left = left.Next
			right.Next = right.Next.Next
		} else {
			// 不小于基准，指针右移
			right = right.Next
		}
	}

	// 已排序链表接在原链表前面
	left.Next = head.Next
	head.Next = dummy.Next
	quickSort(head, pivot)
	quickSort(pivot, tail)
}

func quickSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	quickSort(dummy, nil)
	return dummy.Next
}
