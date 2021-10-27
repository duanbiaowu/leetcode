package leetcode

func numDecodings(s string) int {
	n := len(s)
	// x = f[i-2], y = f[i-1], z = f[i]
	x, y, z := 0, 1, 0
	for i := 1; i <= n; i++ {
		z = 0
		if s[i-1] != '0' {
			z += y
		}
		if i > 1 && s[i-2] != '0' && ((s[i-1]-'0')+(s[i-2]-'0')*10 <= 26) {
			z += x
		}
		x, y = y, z
	}
	return z
}

func numDecodings2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	f := make([]int, n+1)
	f[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}
