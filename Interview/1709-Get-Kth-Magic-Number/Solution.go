package leetcode

func getKthMagicNumber(k int) int {
	if k <= 0 {
		return 0
	}

	res := make([]int, k)
	res[0] = 1
	p3, p5, p7 := 0, 0, 0
	for i := 1; i < k; i++ {
		val := min(min(res[p3]*3, res[p5]*5), res[p7]*7)
		if val%3 == 0 {
			p3++
		}
		if val%5 == 0 {
			p5++
		}
		if val%7 == 0 {
			p7++
		}
		res[i] = val
	}
	return res[k-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
