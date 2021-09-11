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
