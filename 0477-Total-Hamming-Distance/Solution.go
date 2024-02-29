package leetcode

// 暴力方案: 提交运行超时
func totalHammingDistance(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += hammingWeight(uint32(nums[i] ^ nums[j]))
		}
	}

	return res
}

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

// 在暴力方案的基础上进行算法优化
// 优化思路
// 暴力方案的问题在于会计算两个数字的所有比特位是否相同，时间复杂度退化为 O(N^2)
// 但其实只需要计算出两个数字相同 (对应) 的比特位是否相同，而不同比特位之间是否相同无需关心

// 数组 [4, 14, 2] 的距离 =  6
// 其实也就是每个比特位 [0 的数量] * [1 的数量]，最后累加

// 0 1 0 0
// 1 1 1 0
// 0 0 1 0

// 第 0 位的距离为 0
// 第 1 位的距离为 2
// 第 2 位的距离为 2
// 第 3 位的距离为 2
// 0 + 2 + 2 + 2 = 6

// 对于数组中的某个元素 M, 如果其第 N 位为 1
// 那么只需要计算出数组的其他元素中，有多少个元素 第 N 位为 0，也就算出了数组元素 M 和数组的其他元素在第 N 位的距离
// 最后得到如下的计算方式
// 如果长度为 N 的数组，所有元素在 i 位 (0 ~ 30) 上面为 1 有 c 个位置，为 0 有 N - c 个位置，则数组所有元素在 i 位的距离为:
// c * (N - c)
// 从二进制位的最低位到最高位，逐个统计每一位的距离，最后累计即可得到结果
// 对于数组中的每个元素 num, 使用 num >> i & 取出其第 i 位的值
func totalHammingDistanceOpt(nums []int) int {
	var res int
	n := len(nums)

	for i := 0; i < 30; i++ {
		cnt := 0
		for _, num := range nums {
			cnt += num >> i & 1
		}
		res += cnt * (n - cnt)
	}

	return res
}
