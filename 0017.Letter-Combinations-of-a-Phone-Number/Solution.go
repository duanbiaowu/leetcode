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

func backtrackLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res []string
	var str string
	backtrack(digits, &res, str, 0)
	return res
}

func backtrack(digits string, res *[]string, str string, begin int) {
	if begin == len(digits) {
		*res = append(*res, str)
	} else {
		strMap := letterArray[digits[begin]-'2']
		for i := range strMap {
			backtrack(digits, res, str+string(strMap[i]), begin+1)
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

	res := []string{}
	for _, c1 := range letterMap[digits[0]] {
		for _, c2 := range dfsLetterCombinations(digits[1:]) {
			res = append(res, c1+c2)
		}
	}

	return res
}

func bfsLetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	res := letterMap[digits[0]]
	for index := 1; index < len(digits); index++ {
		queue := []string{}

		for _, c1 := range res {
			for _, c2 := range letterMap[digits[index]] {
				queue = append(queue, c1+c2)
			}
		}

		res = queue
	}

	return res
}
