package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connectBFS(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		// 不断将节点入队、出队影响到了执行效率
		for i := 0; i < n; i++ {
			if i+1 < n {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}

	return root
}

func connectBFSOpt(root *Node) *Node {
	if root == nil {
		return nil
	}

	// cur 表示当前所在层
	cur := root
	for cur != nil {
		// 表示下一层的 “哨兵节点”，用于保留下一层链表头节点的指针
		dummy := &Node{}

		// prev 表示下一层节点的遍历指针，指向 “当前节点”
		// 初始化时指向下一层的 “哨兵节点”
		prev := dummy

		// 将下一层的所有节点连接起来，形成一个链表
		// 链表的头节点就是 dummy.Next

		// 每次开始循环之前，cur 都指向当前层的第一个节点
		// 循环执行中，cur 会不断通过 cur.Next 向右移动
		// 本质上和链表的遍历过程是一致的
		for cur != nil {
			// 连接左节点
			if cur.Left != nil {
				prev.Next = cur.Left
				prev = prev.Next
			}
			// 连接右节点
			if cur.Right != nil {
				prev.Next = cur.Right
				prev = prev.Next
			}
			cur = cur.Next
		}

		// 当前层下移一层，转到下一层
		cur = dummy.Next
	}

	return root
}

// 没有 BFS 层级方法直观好理解
// 需要注意 子树==nil 的情况
func connectIteratively(root *Node) *Node {
	var left, prev *Node
	var handle func(root *Node)
	handle = func(root *Node) {
		if root == nil {
			return
		}
		if prev != nil {
			prev.Next = root
		}
		if left == nil {
			left = root
		}
		prev = root
	}

	start := root
	for start != nil {
		left, prev = nil, nil
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = left
	}

	return root
}

// 递归版本
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Left != nil && root.Right == nil {
		root.Left.Next = getNext(root.Next)
	}
	if root.Right != nil {
		root.Right.Next = getNext(root.Next)
	}

	connect(root.Right)
	connect(root.Left)

	return root
}

func getNext(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}
	if root.Next != nil {
		return getNext(root.Next)
	}
	return nil
}
