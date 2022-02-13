package leetcode

func merge(A []int, m int, B []int, n int) {
	k := m + n - 1
	for m > 0 && n > 0 {
		if A[m-1] < B[n-1] {
			A[k] = B[n-1]
			n--
		} else {
			A[k] = A[m-1]
			m--
		}
		k--
	}
	for n > 0 {
		A[k] = B[n-1]
		n--
		k--
	}
}
