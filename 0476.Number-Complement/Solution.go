package leetcode

func findComplement(num int) int {
	res := 0

	for n := num; n > 0; {
		n >>= 1
		res = (res << 1) + 1
	}

	return res ^ num
}
