package leetcode

func constructArr(a []int) []int {
	n := len(a)
	res := make([]int, n)
	for cur, i := 1, 0; i < n; i++ {
		res[i] = cur
		cur *= a[i]
	}
	for cur, i := 1, n-1; i >= 0; i-- {
		res[i] *= cur
		cur *= a[i]
	}
	return res
}
