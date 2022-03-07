package leetcode

func replaceNonCoprimes(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := []int{nums[0]}
	for i := 1; i < n; i++ {
		res = append(res, nums[i])
		for len(res) > 1 {
			x, y := res[len(res)-1], res[len(res)-2]
			g := gcd(x, y)
			if g == 1 {
				break
			}
			res = res[:len(res)-1]
			// 最小公倍数: [6, 4] = 2
			// 6 * 4 / 2 = 12
			// 4 * 6 / 2 = 12
			res[len(res)-1] *= x / g
		}
	}
	return res
}

func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}
