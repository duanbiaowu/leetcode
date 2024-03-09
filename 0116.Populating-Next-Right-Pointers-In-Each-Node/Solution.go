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

	// 使用迭代来实现层级遍历
	for left := root; left.Left != nil; left = left.Left {
		// 遍历当前层，也就是遍历当前层所有节点连接起来形成的 “链表”
		// 将下一层的所有节点连接为一个链表
		for cur := left; cur != nil; cur = cur.Next {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
		}
		// 当前层遍历完成结束后，下一层的所有节点已经连接为一个链表
	}

	return root
}
