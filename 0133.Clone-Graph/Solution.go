package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return cloneGraphDFS(node, make(map[*Node]*Node))
}

func cloneGraphDFS(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}

	if _, ok := visited[node]; ok {
		return visited[node]
	}

	cloneNode := &Node{node.Val, []*Node{}}
	visited[node] = cloneNode

	for _, v := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, cloneGraphDFS(v, visited))
	}
	return cloneNode
}

func cloneGraphBFS(node *Node) *Node {
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
