package leetcode

func integerReplacement(n int) int {
	if n <= 1 {
		return 0
	}
	if n&1 == 0 {
		return 1 + integerReplacement(n>>1)
	}
	return 2 + min(integerReplacement(n>>1), integerReplacement(n>>1+1))
}

var (
	memo = make(map[int]int)
)

func integerReplacementOpt(n int) int {
	if n <= 1 {
		return 0
	}

	v, ok := memo[n]
	if ok {
		return v
	}
	if n&1 == 0 {
		v = 1 + integerReplacementOpt(n>>1)
	} else {
		v = 2 + min(integerReplacementOpt(n>>1), integerReplacementOpt(n>>1+1))
	}

	memo[n] = v
	return v
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
