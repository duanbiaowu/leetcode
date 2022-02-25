package leetcode

// reference: https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/comments/256642

// nums = [1,2,10,4,1,4,3,3]
// a^a=0
// a^0=a
// a^b^c=a^c^b
// a&(-a)=最低位为1的二进制（从右到左）
// 所有的异或结果得到sum=2^10=8
// flag=-8&8=8
// 可分为两组，一组为与flag相与等于1的[10]，另一组为0的[1,2,4,1,4,3,3]
// 组内异或分别得到【10】【2】
func singleNumbers(nums []int) []int {
	sum := 0
	// 得到异或结果，即为不相同两个数的异或结果sum
	for i := range nums {
		sum ^= nums[i]
	}
	if sum == 0 {
		return nil
	}
	// 得到sum的二进制的1的最低位
	flag := sum & (-sum)
	res := []int{0, 0}
	// 分成两个组进行异或，每组异或后的结果就是不相同两个数的其中之一
	for i := range nums {
		if nums[i]&flag == 0 {
			res[0] ^= nums[i]
		} else {
			res[1] ^= nums[i]
		}
	}
	return res
}
