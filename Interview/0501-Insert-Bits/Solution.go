package leetcode

func insertBits(N int, M int, i int, j int) int {
	// N    = 10000000000 (1024)
	// 把最左边的部分调整好了，即抛弃了替换部分和低位部分
	// left = 1000
	left := N >> j >> 1

	// 因此最后要进行或运算，所以把它再移到原来的高位上
	// left = 10000000000
	left = left << j << 1

	// 替换N的 j<-i 位，那么只需要将M左移i位即可
	// mid  = 1001100
	mid := M << i
	// 只需要N的低位，将高位置零,(1<<2)-1 = (11)2
	// right = 0
	right := N & ((1 << i) - 1)
	// left | mid = 10001001100
	// left | mid | right = 10001001100
	return left | mid | right
}
