package leetcode

func singleNumber2(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = (a ^ nums[i]) & ^b
		b = (b ^ nums[i]) & ^a
	}
	return a
}

func singleNumber(nums []int) int {
	// 题目中说明参数为 32 位
	// 创建一个长度为 32 的数组，用于存储所有数字各二进制位中为 1 的次数
	// 本质上也就是通过取模运算来确认结果数字的二进制表示中每个位的值是 0 还是 1
	cnt := [32]int32{}

	for _, num := range nums {
		for j := 0; j < 32; j++ {
			// Go 语言中没有 >>> 运算符
			// 需要先将 num 转为无符号整数进行右移
			cnt[j] += int32(num) >> j & 1
		}
	}

	// 因为除了结果数字之外，其他数字都出现了 3 次
	// 所以对 cnt 数组中各元素进行 % 3 运算，得到 “只出现一次数字” 的各个二进制位
	//    该数字的二进制位只有 2 种结果: 1, 0
	//    1 表示该数字的对应的二进制为 1
	//    0 表示该数字的对应的二进制为 0
	const m = 3
	for i := range cnt {
		cnt[i] %= m
	}

	// 利用位移 + OR 操作，将 cnt 数组中各个二进制位的值恢复到 “只出现一次数字”
	var res int32
	for i := range cnt {
		res <<= 1
		res |= cnt[31-i]
	}

	return int(res)
}

// singleNumber 方法经过简单的优化
func singleNumberOpt(nums []int) int {
	cnt := [32]int32{}

	for _, num := range nums {
		for j := 0; j < 32; j++ {
			cnt[j] += int32(num) >> j & 1
		}
	}

	const m = 3
	var res int32

	for i := range cnt {
		res <<= 1
		res |= cnt[31-i] % m
	}

	return int(res)
}

// 在 singleNumberOpt 方法的基础上可以继续优化
// 我们可以计算出 “只出现一个数字” 的各个二进制位
// 然后在循环中直接通过 OR 运算进行位移操作
func singleNumberOpt2(nums []int) int {
	const m = 3
	res := int32(0)

	for i := 0; i < 32; i++ {
		var sum int32
		for _, num := range nums {
			sum += int32(num) >> i & 1
		}

		res |= (sum % m) << i
	}

	return int(res)
}
