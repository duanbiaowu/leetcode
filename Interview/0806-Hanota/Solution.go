package leetcode

func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}
	move(len(A), &A, &B, &C)
	return C
}

// n = 1 时，直接把盘子从 A 移到 C
// n > 1 时
// 	先把上面 n - 1 个盘子从 A 移到 B（子问题，递归）
//	再将最大的盘子从 A 移到 C
//	再将 B 上 n - 1 个盘子从 B 移到 C（子问题，递归）
func move(n int, A, B, C *[]int) {
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
	} else {
		// 将 A 上面 n-1 个通过空的 C 移到 B
		move(n-1, A, C, B)
		// 将 A 最后一个移到 C
		move(1, A, B, C)
		// 将 B 上面 n-1 个通过空的 A 移到 C
		move(n-1, B, A, C)
	}
}
