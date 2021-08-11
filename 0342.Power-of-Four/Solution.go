package leetcode

func isPowerOfFour(n int) bool {
	// 数论
	//return n > 0 && n&(n-1) == 0 && n%3 == 1

	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
