package leetcode

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	// 确保负数以 32 位二进制补码的形式表示
	num &= 0xFFFFFFFF

	hex := "0123456789abcdef"
	var res string
	for num != 0 {
		res = string(hex[num&0xf]) + res
		num >>= 4
	}

	return res
}
