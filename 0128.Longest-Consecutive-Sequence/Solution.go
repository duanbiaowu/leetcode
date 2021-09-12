package leetcode

func longestConsecutive(nums []int) int {
	set := map[int]int{}
	res := 0

	for _, v := range nums {
		if _, ok := set[v]; !ok {
			left, leftOK := set[v-1]
			if !leftOK {
				left = 0
			}
			right, rightOK := set[v+1]
			if !rightOK {
				right = 0
			}

			cur := left + right + 1
			if cur > res {
				res = cur
			}

			set[v] = cur
			set[v-left] = cur
			set[v+right] = cur
		}
	}

	return res
}

func longestConsecutive2(nums []int) int {
	set := map[int]struct{}{}
	for _, v := range nums {
		set[v] = struct{}{}
	}

	res := 0
	for v := range set {
		if _, ok := set[v-1]; !ok {
			num := v
			step := 1
			for _, ok = set[num+1]; ok; _, ok = set[num+1] {
				num++
				step++
			}
			if step > res {
				res = step
			}
		}
	}
	return res
}
