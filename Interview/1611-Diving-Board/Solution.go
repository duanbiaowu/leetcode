package leetcode

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}

	var res []int
	if shorter == longer {
		res = append(res, shorter*k)
		return res
	}

	// 一开始全部为 shorter 保证最小
	cnt := shorter * k
	res = append(res, cnt)
	for i := k; i > 0; i-- {
		// 减 1 个 shorter，加 1 个 longer
		cnt = cnt - shorter + longer
		res = append(res, cnt)
	}
	return res
}
