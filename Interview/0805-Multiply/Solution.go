package leetcode

// 13 * 12
// = 13 * (8 + 4)
// = 13 * 8 + 13 * 4
// = (13 << 3) + (13 << 2)
func multiply(A int, B int) int {
	var max, min int
	if A > B {
		max = A
		min = B
	} else {
		max = B
		min = A
	}

	res := 0
	for i := 0; min != 0; i++ {
		if min&1 == 1 {
			res += max << i
		}
		min >>= 1
	}
	return res
}
