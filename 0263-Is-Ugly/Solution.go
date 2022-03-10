package leetcode

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	divisors := [3]int{2, 3, 5}
	for i := range divisors {
		for n%divisors[i] == 0 {
			n /= divisors[i]
		}
	}
	return n == 1
}
