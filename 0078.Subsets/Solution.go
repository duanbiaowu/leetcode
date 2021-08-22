package leetcode

// a = {5, 2, 9}
// 0/1 序列对应的二进制数正好从 0 到 2^n - 1
// ---0/1序列------------子集--------------0/1序列对应的二进制数------------
/**
      000               {}                0
 	  001				{9}				  1
	  010				{2}				  2
      011				{2, 9}			  3
      100				{5}				  4
      101				{5, 9}			  5
      110				{5, 2}			  6
      111				{5, 2, 9}		  7
*/
func subsets(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		res = append(res, append([]int{}, set...))
	}
	return res
}
