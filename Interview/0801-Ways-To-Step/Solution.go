package leetcode

// 一次可以上1阶、2阶或3阶台阶
func waysToStep(n int) int {
	if n <= 0 {
		return 0
	}
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}
	const MOD = 1000000007
	step1, step2, step3 := 1, 2, 4
	cnt := 0
	for i := 4; i <= n; i++ {
		cnt = ((step1+step2)%MOD + step3) % MOD
		step1 = step2
		step2 = step3
		step3 = cnt
	}
	return cnt
}
