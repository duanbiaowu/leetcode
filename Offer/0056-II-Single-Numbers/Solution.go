package leetcode

// reference: https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/mian-shi-ti-56-ii-shu-zu-zhong-shu-zi-chu-xian-d-4/
func singleNumber(nums []int) int {
	// 长度 32 的数组记录所有数字的各二进制位的 1 的出现次数
	cnt := [32]int{}
	for i := range nums {
		for j := 0; j < 32; j++ {
			// 更新第 j 位
			cnt[j] += nums[i] & 1
			// 第 j 位 --> 第 j + 1 位
			nums[i] >>= 1
		}
	}
	// 修改 mod ，实现除了一个数字以外，其余数字都出现 m 次的通用问题
	res, mod := 0, 3
	for i := 0; i < 32; i++ {
		res <<= 1
		// 恢复第 i 位的值到 res
		res |= cnt[31-i] % mod
	}
	return res
}
