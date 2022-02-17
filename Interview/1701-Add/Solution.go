package leetcode

// 异或: 相同 = 0， 不同 = 1
func add(a int, b int) int {
	sum, carry := 0, 0
	for b != 0 {
		// 异或计算未进位的部分
		sum = a ^ b
		// 进位部分
		carry = (a & b) << 1
		// 保存未进位部分，再次计算
		a = sum
		// 保存进位部分，再次计算
		b = carry
	}
	return a
}

func add2(a int, b int) int {
	if b == 0 {
		return a
	}
	return add2(a^b, (a&b)<<1)
}
