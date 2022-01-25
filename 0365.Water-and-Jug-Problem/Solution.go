package leetcode

func canMeasureWater(x int, y int, z int) bool {
	stack := [][2]int{
		{0, 0},
	}
	seen := make(map[int64]bool)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		key := hash(top)
		if _, ok := seen[key]; ok {
			stack = stack[:len(stack)-1]
			continue
		}

		seen[key] = true
		remainX, remainY := top[0], top[1]
		if remainX == z || remainY == z || remainX+remainY == z {
			return true
		}

		// 把 X 壶装满
		stack = append(stack, [2]int{x, remainY})
		// 把 Y 壶装满
		stack = append(stack, [2]int{remainX, y})
		// 把 X 壶倒空
		stack = append(stack, [2]int{0, remainY})
		// 把 Y 壶倒空
		stack = append(stack, [2]int{remainX, 0})
		// 把 X 壶的水倒入 Y 壶，直至倒满或倒空
		stack = append(stack, [2]int{
			remainX - min(remainX, y-remainY),
			remainY + min(remainX, y-remainY),
		})
		// 把 Y 壶的水倒入 X 壶，直至倒满或倒空
		stack = append(stack, [2]int{
			remainX + min(remainY, x-remainX),
			remainY - min(remainY, x-remainX),
		})
	}

	return false
}

func hash(state [2]int) int64 {
	return int64(state[0]*1000001 + state[1])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
