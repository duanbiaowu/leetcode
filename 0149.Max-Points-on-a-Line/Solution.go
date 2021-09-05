package leetcode

type K struct {
	m int
	n int
}

// https://leetcode-cn.com/problems/max-points-on-a-line/solution/zhi-xian-shang-zui-duo-de-dian-shu-by-le-tq8f/
// 四个小优化:
// 1. 在点的总数量小于等于 2 的情况下，可以用一条直线将所有点串联，直接返回点的总数量即可
// 2. 枚举到点 i 时，只需要考虑编号大于 i 的点到点 i 的斜率，
//	  因为如果直线同时经过编号小于点 i 的点 j，当我们枚举到 j 时就已经考虑过该直线了
// 3. 找到一条直线经过了图中超过半数的点时，可以确定该直线即为经过最多点的直线
// 4. 枚举到点 i（假设编号从 0 开始）时，我们至多只能找到 n-i 个点共线
//	  假设此前找到的共线的点的数量的最大值为 k，如果有 k >= n−i
//	  那么此时我们即可停止枚举，因为不可能再找到更大的答案了
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	var res int
	for i := 0; i < n; i++ {
		if res >= n-i || res > n>>1 {
			break
		}

		counter := map[K]int{}
		for j := i + 1; j < n; j++ {
			k := getK(points[i], points[j])
			counter[k]++
		}

		for _, c := range counter {
			res = max(res, c+1)
		}
	}

	return res
}

// 根据两点求斜率
func getK(p1, p2 []int) K {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	// p1,p2是同一点
	if dx == 0 && dy == 0 {
		return K{0, 0}
	}
	// p1,p2斜率为无穷
	if dx == 0 {
		return K{1, 0}
	}
	// p1,p2斜率为0
	if dy == 0 {
		return K{0, 1}
	}
	// 改变x, y方向
	if dx < 0 {
		dx = -dx
		dy = -dy
	}
	d := gcd(dx, dy)
	// 返回最简约之比
	return K{dx / d, dy / d}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
