package leetcode

func isPalindrome(x int) bool {
	// 当x为负数, 10的倍数时不是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// 利用回文数性质, 反转一半数字
	// 例: 1221, 前半部分 12 与 后半部分反转后 (21 => 12) 进行比较
	halfReverseX := 0
	for x > halfReverseX {
		halfReverseX = halfReverseX*10 + x%10
		x /= 10
	}

	// 需要覆盖个位数
	return x == halfReverseX || halfReverseX/10 == x
}
