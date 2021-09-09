package leetcode

func candy(ratings []int) (ans int) {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	res, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			res += pre
			inc = pre
		} else {
			// 逆向操作: 画图更直观 => [1, 3, 5, 3, 2, 1]
			dec++
			if dec == inc {
				dec++
			}
			res += dec
			pre = 1
		}
	}

	return res
}

func candy2(ratings []int) (ans int) {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	left := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	res, right := 0, 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(right, left[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
