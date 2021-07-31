package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNodesByArray(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}