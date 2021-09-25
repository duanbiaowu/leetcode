package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
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

// 没有 BFS 层级方法直观好理解
// 需要注意 子树==nil 的情况
func connect2(root *Node) *Node {
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
