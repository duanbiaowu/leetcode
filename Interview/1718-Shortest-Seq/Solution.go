package leetcode

// 滑动窗口
func shortestSeq(big []int, small []int) []int {
	if len(big) == 0 || len(small) == 0 {
		return nil
	}

	var res []int
	m := make(map[int]int)
	count, start := 0, 0

	for _, v := range small {
		if _, ok := m[v]; !ok {
			count++
		}
		m[v]++
	}

	for end := 0; end < len(big); end++ {
		m[big[end]]--
		if m[big[end]] == 0 {
			count--
		}
		for count == 0 {
			m[big[start]]++
			if m[big[start]] > 0 {
				count++
				if len(res) == 0 || res[1]-res[0] > end-start {
					res = []int{start, end}
				}
			}
			start++
		}
	}

	return res
}
