package leetcode

func climbStairs(n int) int {
	s1, s2, s3 := 1, 1, 1

	for i := 1; i < n; i++ {
		s1 = s2
		s2 = s3
		s3 = s1 + s2
	}
	return s3
}

// timeout
func climbStairsRecursively(n int) int {
	if n < 3 {
		return n
	}
	return climbStairsRecursively(n-1) + climbStairsRecursively(n-2)
}
