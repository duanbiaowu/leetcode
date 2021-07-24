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

// backtrack
func backtrackLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	var str string
	backtrack(digits, &res, str, 0)
	return res
}

func backtrack(digits string, res *[]string, str string, start int) {
	if start == len(digits) {
		*res = append(*res, str)
	} else {
		strMap := letterMap[digits[start]-'2']
		for i := 0; i < len(strMap); i++ {
			backtrack(digits, res, str+string(strMap[i]), start+1)
		}
	}
}