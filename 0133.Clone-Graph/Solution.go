package leetcode

// Node Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return dfs(node, make(map[*Node]*Node))
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}

	if _, ok := visited[node]; ok {
		return visited[node]
	}

	// 克隆当前节点
	cloneNode := &Node{node.Val, []*Node{}}
	// 将当前节点标记为已经克隆
	visited[node] = cloneNode

	// 遍历当前的节点的邻节点，递归克隆
	for _, v := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(v, visited))
	}

	return cloneNode
}

func bfs(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[*Node]*Node)
	visited[node] = &Node{node.Val, []*Node{}}
	queue := []*Node{node}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for _, v := range top.Neighbors {
			if _, ok := visited[v]; !ok {
				visited[v] = &Node{v.Val, []*Node{}}
				queue = append(queue, v)
			}
			visited[top].Neighbors = append(visited[top].Neighbors, visited[v])
		}
	}

	return visited[node]
}
