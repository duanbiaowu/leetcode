package leetcode

// 通用算法: 所有小于等于 n 的非负整数中数字 X 出现的个数
// 带入数据(25)跟踪一下，理解过程
func numberOf2sInRange(n int) int {
	X := 2
	count := 0
	for base := 1; base <= n; base *= 10 {
		leftPart, rightPart := n/base, n%base
		count += leftPart / 10 * base
		if lastBit := leftPart % 10; lastBit > X {
			count += base
		} else if lastBit == X {
			count += rightPart + 1
		}
	}
	return count
}
