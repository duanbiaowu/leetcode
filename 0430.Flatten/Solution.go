package leetcode

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Example:
//
//		1---2---3---4---5---6--NULL
//	            |
//			    7---8---9---10--NULL
//					|
//					11--12--NULL
func flatten(root *Node) *Node {
	cur := root

	for cur != nil {
		if cur.Child != nil {
			next := cur.Next
			node := flatten(cur.Child)
			cur.Child = nil

			cur.Next = node
			node.Prev = cur

			for cur.Next != nil {
				cur = cur.Next
			}
			if next != nil {
				cur.Next = next
				next.Prev = cur
			}
		}

		cur = cur.Next
	}

	return root
}
