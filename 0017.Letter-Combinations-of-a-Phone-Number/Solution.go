package leetcode

var letterArray = []string{
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
		strMap := letterArray[digits[start]-'2']
		for i := 0; i < len(strMap); i++ {
			backtrack(digits, res, str+string(strMap[i]), start+1)
		}
	}
}

var letterMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func dfsLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return letterMap[digits[0]]
	}
	var res []string
	for _, c1 := range letterMap[digits[0]] {
		for _, c2 := range dfsLetterCombinations(digits[1:]) {
			res = append(res, c1+c2)
		}
	}
	return res
}
