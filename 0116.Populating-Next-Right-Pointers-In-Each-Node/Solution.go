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
