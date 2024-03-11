package leetcode

func isHappy(n int) bool {
	m := make(map[int]struct{})

	for n != 1 {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		} else {
			break
		}
		n = getNext(n)
	}

	return n == 1
}

func getNext(n int) int {
	sum := 0

	for n > 0 {
		val := n % 10
		sum += val * val
		n /= 10
	}

	return sum
}

func isHappy2(n int) bool {
	slow, fast := n, getNext(n)
	for fast != 1 && fast != slow {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}
