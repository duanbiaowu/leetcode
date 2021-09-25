package leetcode

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

// 直接使用组合公式C(n,i) = n! / (i! * (n-i)! )
// 第(i+1)项是第i项的倍数= i * (n-i) / (i+1)
func getRow2(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 0; i < rowIndex; i++ {
		res[i+1] = res[i] * (rowIndex - i) / (i + 1)
	}
	return res
}
