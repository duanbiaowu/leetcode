package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	root.Left.Next = root.Right
	// 限定为完全二叉树
	//if root.Next != nil {

	// 不限定为完全二叉树
	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)

	return root
}

func connectIteratively(root *Node) *Node {
	if root == nil {
		return root
	}

	for left := root; left.Left != nil; left = left.Left {
		for cur := left; cur != nil; cur = cur.Next {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
		}
	}

	return root
}
