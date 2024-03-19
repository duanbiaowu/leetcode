package leetcode

type edge struct {
	to     int
	weight float64
}

func buildMap(equations [][]string) map[string]int {
	idMap := make(map[string]int)
	for _, eq := range equations {
		x, y := eq[0], eq[1]
		if _, ok := idMap[x]; !ok {
			idMap[x] = len(idMap)
		}
		if _, ok := idMap[y]; !ok {
			idMap[y] = len(idMap)
		}
	}

	return idMap
}

// 根据变量数组和值数组构建图
// 两个点对象之间计算好对应的权重
func buildGraph(equations [][]string, values []float64, idMap map[string]int) [][]*edge {
	graph := make([][]*edge, len(idMap))

	for i, eq := range equations {
		v, w := idMap[eq[0]], idMap[eq[1]]
		graph[v] = append(graph[v], &edge{
			to:     w,
			weight: values[i],
		})
		graph[w] = append(graph[w], &edge{
			to:     v,
			weight: 1 / values[i],
		})
	}

	return graph
}

func bfs(graph [][]*edge, start, end int) float64 {
	ratios := make([]float64, len(graph))
	// 每个值到自身的权重为 1, 例如 2/2 = 1
	ratios[start] = 1
	queue := []int{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == end {
			return ratios[v]
		}

		for _, e := range graph[v] {
			if w := e.to; ratios[w] == 0 {
				ratios[w] = ratios[v] * e.weight
				queue = append(queue, w)
			}
		}
	}

	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	idMap := buildMap(equations)
	graph := buildGraph(equations, values, idMap)

	res := make([]float64, len(queries))
	for i, query := range queries {
		x, y := query[0], query[1]
		start, hasS := idMap[x]
		end, hasE := idMap[y]

		// 如果一个元素不存在，说明两者之间没有连通路径
		// 距离直接置为 -1 即可
		if !hasS || !hasE {
			res[i] = -1
		} else {
			res[i] = bfs(graph, start, end)
		}
	}

	return res
}

// 根据变量数组和值数组构建图
// 预先计算出任意两点之间的距离
func buildGraphOpt(equations [][]string, values []float64, idMap map[string]int) [][]float64 {
	graph := make([][]float64, len(idMap))

	for i := range graph {
		graph[i] = make([]float64, len(idMap))
	}

	for i, eq := range equations {
		v, w := idMap[eq[0]], idMap[eq[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1 / values[i]
	}

	return graph
}

// 使用 Floyd 算法，预先计算出任意两点之间的距离
func floyd(graph [][]float64) {
	for k := range graph {
		for i := range graph {
			for j := range graph {
				// 测试用例卡精度
				if graph[i][k] > 1e-6 && graph[k][j] > 1e-6 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}
}

func calcEquationOpt(equations [][]string, values []float64, queries [][]string) []float64 {
	idMap := buildMap(equations)
	graph := buildGraphOpt(equations, values, idMap)

	floyd(graph)

	res := make([]float64, len(queries))

	for i, query := range queries {
		x, y := query[0], query[1]
		start, hasS := idMap[x]
		end, hasE := idMap[y]

		if !hasS || !hasE || graph[start][end] == 0 {
			res[i] = -1
		} else {
			res[i] = graph[start][end]
		}
	}

	return res
}
