package leetcode

func replaceSpace(s string) string {
	n := len(s)
	res := make([]byte, n*3)
	index := 0
	for _, v := range s {
		if v == ' ' {
			res[index] = '%'
			res[index+1] = '2'
			res[index+2] = '0'
			index += 3
		} else {
			res[index] = byte(v)
			index++
		}
	}
	return string(res[:index])
}
