package leetcode

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	// a^-x = 1/a^x
	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n>>1)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	// 注意这里不能写为: n>>1
	// -3>>1 = -2
	// -3/2 = -1
	half := myPow2(x, n/2)
	mod := myPow2(x, n%2)
	return half * half * mod
}
