package leetcode

func isNumber(s string) bool {
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	return numFlag
}

type State int
type CharType int

const (
	StateInitial State = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateEx
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharIllegal
)

func toCharType(char byte) CharType {
	switch char {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	default:
		return CharIllegal
	}
}

// DFA
// https://leetcode-cn.com/problems/valid-number/solution/you-xiao-shu-zi-by-leetcode-solution-298l/
func isNumber2(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: {
			CharNumber: StateInteger,
			CharExp:    StateEx,
			CharPoint:  StatePoint,
		},
		StatePoint: {
			CharNumber: StateFraction,
			CharExp:    StateEx,
		},
		StatePointWithoutInt: {
			CharNumber: StateFraction,
		},
		StateFraction: {
			CharNumber: StateFraction,
			CharExp:    StateEx,
		},
		StateEx: {
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: {
			CharNumber: StateExpNumber,
		},
		StateExpNumber: {
			CharNumber: StateExpNumber,
		},
	}

	state := StateInitial
	for i := 0; i < len(s); i++ {
		charType := toCharType(s[i])
		if _, ok := transfer[state][charType]; !ok {
			return false
		} else {
			state = transfer[state][charType]
		}
	}

	return state == StateInteger || state == StatePoint ||
		state == StateFraction || state == StateExpNumber ||
		state == StateEnd
}
