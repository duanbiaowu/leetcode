package leetcode

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return nil
	}

	var ip []int
	var res []string
	backtrack(s, 0, ip, &res)
	return res
}

func backtrack(s string, begin int, ip []int, res *[]string) {
	if begin == len(s) {
		if len(ip) == 4 {
			*res = append(*res, toString(ip))
		}
		return
	}

	if begin == 0 {
		num, _ := strconv.Atoi(string(s[0]))
		ip = append(ip, num)
		backtrack(s, begin+1, ip, res)
	} else {
		num, _ := strconv.Atoi(string(s[begin]))
		cur := ip[len(ip)-1]
		next := cur*10 + num

		if next <= 255 && cur != 0 {
			ip[len(ip)-1] = next
			backtrack(s, begin+1, ip, res)
			ip[len(ip)-1] /= 10
		}

		if len(ip) < 4 {
			ip = append(ip, num)
			backtrack(s, begin+1, ip, res)
			ip = ip[:len(ip)-1]
		}
	}
}

func toString(ip []int) string {
	if len(ip) == 0 {
		return ""
	}

	res := strconv.Itoa(ip[0])
	for i := 1; i < len(ip); i++ {
		res += "." + strconv.Itoa(ip[i])
	}
	return res
}
