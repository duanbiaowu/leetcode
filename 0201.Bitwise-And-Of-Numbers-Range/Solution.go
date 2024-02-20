package leetcode

func rangeBitwiseAnd(left int, right int) int {
	shift := 0

	for left != right {
		left >>= 1
		right >>= 1
		shift++
	}

	return left << shift
}

func rangeBitwiseAndOpt(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}
