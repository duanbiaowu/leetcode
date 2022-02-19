package leetcode

import "fmt"

// 思维转变一下，不是把两个数组比较，看有没有重复的。
// 而是每次的数字出现，都去检查一下是否出现过，在哪个数组里出现的。
// reference: https://leetcode-cn.com/problems/sparse-similarity-lcci/solution/chao-xiang-xi-ti-jie-qing-song-kan-dong-shuang-bai/
func computeSimilarities(docs [][]int) []string {
	n := len(docs)
	m := make(map[int][]int)
	help := make([][]int, n)
	for i := 0; i < n; i++ {
		help[i] = make([]int, n)
	}

	for i := range docs {
		for j := range docs[i] {
			val := docs[i][j]
			if _, ok := m[val]; ok {
				for _, v := range m[val] {
					help[v][i] += 1
				}
				m[val] = append(m[val], i)
			} else {
				m[val] = []int{i}
			}
		}
	}

	var res []string
	for i := range help {
		for j := i + 1; j < n; j++ {
			if help[i][j] > 0 {
				// 合集 - 交集 = 并集
				// 交集 / 并集 = 相似度
				rate := float64(help[i][j]) / float64(len(docs[i])+len(docs[j])-help[i][j])
				res = append(res, fmt.Sprintf("%d,%d: %.4f", i, j, rate+1e-9))
			}
		}
	}
	return res
}
