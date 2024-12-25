package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
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

func GenerateRandomListNodesBySlice(nums [][2]int) *RandomListNode {
	dummy := &RandomListNode{}
	cur := dummy
	m := map[int]*RandomListNode{}

	for _, v := range nums {
		node := &RandomListNode{Val: v[0]}
		m[node.Val] = node
		cur.Next = node
		cur = cur.Next
	}

	for _, v := range nums {
		if node, ok := m[v[1]]; ok {
			m[v[0]].Random = node
		}
	}

	return dummy.Next
}

func DumpListNode(head *ListNode) {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	fmt.Println(res)
}
