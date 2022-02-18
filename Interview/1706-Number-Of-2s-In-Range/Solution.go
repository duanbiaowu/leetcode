package leetcode

// 带入数据(25)跟踪一下，理解过程
func numberOf2sInRange(n int) int {
	count := 0
	for base := 1; base <= n; base *= 10 {
		leftPart, rightPart := n/base, n%base
		count += leftPart / 10 * base
		if lastBit := leftPart % 10; lastBit > 2 {
			count += base
		} else if lastBit == 2 {
			count += rightPart + 1
		}
	}
	return count
}
