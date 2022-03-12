package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*Node{root}
	visited := map[*Node]bool{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		// 如果当前节点为叶子节点或者当前节点的子节点已经遍历过
		if len(node.Children) == 0 || visited[node] {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
		} else {
			// 从右向左入栈，出栈时最左边节点先出
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			visited[node] = true
		}
	}

	return res
}

func postorder2(root *Node) []int {
	var res []int
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for i := range root.Children {
		helper(root.Children[i], res)
	}
	*res = append(*res, root.Val)
}
