package leetcode

func binaryGap(n int) int {
	var res int

	//           1 0 1 1 0
	//  	     		  	<-left <- right
	// left, right 相当于两个前后两个指针，指向不同的位置
	// right 变量初始化为 -1 而不是 0
	// 因为可能会存在两种边界情况:
	//   1. n 的比特位全部为 1，例如 7 = 111
	//   2. n 的第一个比特位为 1，这种情况下如果 right 初始化为 0，判断条件就会变为: right > 0, 那么就会漏掉一个计数
	for right, left := -1, 0; n != 0; left++ {
		if n&1 == 1 {
			if right > -1 {
				res = max(res, left-right)
			}
			right = left
		}
		n >>= 1
	}

	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
