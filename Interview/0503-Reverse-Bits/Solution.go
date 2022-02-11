package leetcode

func reverseBits(num int) int {
	max, pre, cur := 0, 0, 0
	for bits := 32; bits > 0; bits-- {
		// 因为只能翻转 1 次，所以再次翻转时 cur - pre
		if num&1 == 0 {
			cur -= pre
			pre = cur + 1
		}
		cur++
		if cur > max {
			max = cur
		}
		num >>= 1
	}
	return max
}
