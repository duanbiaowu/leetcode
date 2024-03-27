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
	// 边界处理
	if node == nil {
		return node
	}

	// 声明节点访问标记 Map
	// 除了访问标记外，该 Map 还有另外一个作用
	// 那就是将源节点和克隆节点进行映射
	visited := make(map[*Node]*Node)
	// 克隆起始节点 (将起始节点标记为已访问)
	visited[node] = &Node{node.Val, []*Node{}}

	// 将起始节点添加到队列中
	queue := []*Node{node}

	for len(queue) > 0 {
		// 从队列中取出第一个节点
		top := queue[0]
		// 第一个节点出队
		queue = queue[1:]

		for _, v := range top.Neighbors {
			if _, ok := visited[v]; !ok {
				// 如果邻居节点还未被克隆过 (未被访问过)
				// 克隆邻居节点 (标记为已访问)
				visited[v] = &Node{v.Val, []*Node{}}

				// 将邻居节点添加到队列中
				queue = append(queue, v)
			}

			// 更新当前节点的邻居节点
			visited[top].Neighbors = append(visited[top].Neighbors, visited[v])
		}
	}

	// 返回起始节点对应的克隆新节点
	return visited[node]
}
