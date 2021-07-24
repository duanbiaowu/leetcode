package leetcode

var letterMap = []string{
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

// DFS
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	var str string
	dfs(digits, &res, str, 0)
	return res
}

func dfs(digits string, res *[]string, str string, start int) {
	if start == len(digits) {
		*res = append(*res, str)
	} else {
		strMap := letterMap[digits[start]-'2']
		for i := 0; i < len(strMap); i++ {
			dfs(digits, res, str+string(strMap[i]), start+1)
		}
	}
}
