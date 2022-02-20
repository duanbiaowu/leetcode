package leetcode

func getValidT9Words(num string, words []string) []string {
	keys := [26]byte{
		'2', '2', '2',
		'3', '3', '3',
		'4', '4', '4',
		'5', '5', '5',
		'6', '6', '6',
		'7', '7', '7', '7',
		'8', '8', '8',
		'9', '9', '9', '9',
	}

	var res []string
	n := len(num)
	for i := range words {
		if len(words[i]) != n {
			continue
		}
		flag := true
		for j := range words[i] {
			if num[j] != keys[words[i][j]-'a'] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, words[i])
		}
	}
	return res
}
