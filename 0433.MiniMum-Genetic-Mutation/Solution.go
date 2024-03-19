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
	for step := 0; len(queue) > 0; step++ {
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
							return step + 1
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
