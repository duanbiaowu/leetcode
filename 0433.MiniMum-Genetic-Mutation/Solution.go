package leetcode

func minMutation(start string, end string, bank []string) int {
	// 剪枝 1
	if start == end {
		return 0
	}

	set := make(map[string]struct{})
	for _, s := range bank {
		set[s] = struct{}{}
	}
	// 剪枝 2
	if _, ok := set[end]; !ok {
		return -1
	}

	// 变化字符只能从 ‘A’, ‘C’, ‘G’, ‘T’ 中进行选择
	const switchChars = "ACGT"

	queue := []string{start}
	// 变化一次最多可能会生成 3 × 8 = 24 种不同的基因序列
	// 一个序列为 8 个字符，每个字符变为其他 3 个字符
	// 虽然有 24 种不同的基因序列，但是改变次数只计 1 次
	for step := 1; len(queue) > 0; step++ {
		length := len(queue)

		for i := 0; i < length; i++ {
			top := queue[i]
			for j, c := range top {
				for _, sc := range switchChars {
					// 当前字符和切换字符相同时, 直接跳过
					if c == sc {
						continue
					}

					// 将字符的前半部分 + 当前字符改变 + 将字符的后半部分, 拼接为一个新的字符串
					next := top[:j] + string(sc) + top[j+1:]
					if _, ok := set[next]; ok {
						if next == end {
							return step
						}

						// 如果改变后的字符之前已经遍历过，删除记录，避免重复遍历
						delete(set, next)
						queue = append(queue, next)
					}
				}
			}
		}

		queue = queue[length:]
	}

	return -1
}

// 优化
// 可以对 bank 进行预处理，只在合法的基因变化规则中进行搜索即可
// 使用邻接表保存基因库中的每个基因的合法变换
// BFS 搜索过程中，每次只在临界表中搜索即可 (因为不在临界表中的变化不符合规则，是无效的)
func minMutationOpt(start string, end string, bank []string) int {
	// 剪枝 1
	if start == end {
		return 0
	}

	// 使用 邻接表 表示基因合法变化的关系矩阵
	n := len(bank)
	adjacency := make([][]int, n)
	// 记录目标字符串的索引位置
	endIndex := -1

	// 初始化 邻接表
	for i := range bank {
		if bank[i] == end {
			// 更新目标字符串的索引位置
			endIndex = i
		}
		for j := i + 1; j < n; j++ {
			if diffOne(bank[i], bank[j]) {
				adjacency[i] = append(adjacency[i], j)
				adjacency[j] = append(adjacency[j], i)
			}
		}
	}
	// 剪枝 2 (end 字符不存在于 bank 中)
	if endIndex == -1 {
		return -1
	}

	var queue []int
	visited := make([]bool, n)
	for i := range bank {
		// 将可以直接由 start 进行变化的字符串加入到队列中
		if diffOne(start, bank[i]) {
			queue = append(queue, i)
			visited[i] = true
		}
	}

	// BFS 搜索
	// 每次搜索过程算 ”一轮“, 这一轮中需要变化的次数 = 队列的长度
	// 变化次数计数器 + 1
	for step := 1; len(queue) > 0; step++ {
		length := len(queue)

		for i := 0; i < length; i++ {
			top := queue[i]
			if top == endIndex {
				return step
			}

			for _, next := range adjacency[top] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}

		queue = queue[length:]
	}

	return -1
}

// 检测两个字符串中是否只有一个字符不同
func diffOne(s, t string) bool {
	var diff bool

	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}

	return true
}
