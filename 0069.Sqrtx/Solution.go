package leetcode

func mySqrt(x int) int {
	n := x

	for n*n > x {
		n = (n + x/n) >> 1
	}

	return n
}
