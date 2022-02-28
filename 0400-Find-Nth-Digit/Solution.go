package leetcode

import (
	"strconv"
)

// reference: https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/
// 一位数：1~9，共 9 位 = 9 * 1
// 两位数：10~99，共 180 位 = 90 * 2
// 三位数：100~999，共 2700 位 = 900 * 3
func findNthDigit(n int) int {
	digit := 1
	start := 1
	count := 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	str := strconv.Itoa(start + (n-1)/digit)
	return int(str[(n-1)%digit] - '0')
}
