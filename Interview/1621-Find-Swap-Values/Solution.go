package leetcode

func findSwapValues(array1 []int, array2 []int) []int {
	// a - x + y = b - y + x
	// a - b = 2x - 2y
	// a - b = delta

	// delta / 2 = x - y
	// delta / 2 + y = x
	// 差值是偶数

	sum1, sum2 := 0, 0
	for i := 0; i < len(array1); i++ {
		sum1 += array1[i]
	}
	for i := 0; i < len(array2); i++ {
		sum2 += array2[i]
	}
	delta := sum1 - sum2
	if delta&1 == 1 {
		return nil
	}

	// x = delta / 2 + y
	// y = x - delta / 2
	// delta / 2 - x = -y
	delta /= 2
	set := make(map[int]struct{})
	for i := 0; i < len(array1); i++ {
		set[delta-array1[i]] = struct{}{}
	}
	// 根据 -y 寻找 delta / 2 - x
	for i := 0; i < len(array2); i++ {
		y := 0 - array2[i]
		if _, ok := set[y]; ok {
			return []int{delta - y, -y}
		}
	}

	return nil
}
