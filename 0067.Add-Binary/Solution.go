package leetcode

func addBinary(a string, b string) string {
	var res string
	aLen, bLen := len(a), len(b)
	n := max(aLen, bLen)

	carry := 0
	for i := 0; i < n; i++ {
		if i < aLen {
			carry += int(a[aLen-i-1] - '0')
		}
		if i < bLen {
			carry += int(b[bLen-i-1] - '0')
		}
		// 注意: 转换规则
		res = string(byte(carry&1+'0')) + res
		carry >>= 1
	}

	if carry > 0 {
		res = "1" + res
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
