package leetcode

func evenOddBit(n int) []int {
	even, odd := 0, 0

	for i := 0; n != 0; i++ {
		if n&1 == 1 {
			if i&1 == 0 {
				even++
			} else {
				odd++
			}
		}

		n >>= 1
	}

	return []int{even, odd}
}
