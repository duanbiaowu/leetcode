package leetcode

func convertInteger(A int, B int) int {
	x := int32(A ^ B)
	cnt := 0
	for x != 0 {
		x &= x - 1
		cnt++
	}
	return cnt
}
