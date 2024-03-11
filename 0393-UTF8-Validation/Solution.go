package leetcode

func validUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}

	// 用于判断 data 首字节之后一共有多少字节
	bitIdx := 0

	for _, num := range data {
		if bitIdx == 0 {
			if num>>5 == 0b110 {
				bitIdx = 1
			} else if num>>4 == 0b1110 {
				bitIdx = 2
			} else if num>>3 == 0b11110 {
				bitIdx = 3
			} else if num>>7 > 0 {
				return false
			}
		} else {
			// 检查每个 data 首字节后面的几个字节是否合法
			// 后面字节的前两位一律设为 10
			if num>>6 != 0b10 {
				return false
			}
			bitIdx--
		}
	}

	return bitIdx == 0
}
