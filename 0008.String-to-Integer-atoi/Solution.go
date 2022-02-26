package leetcode

import "math"

func myAtoi(s string) int {
	left, right := 0, len(s)

	for left < right && s[left] == ' ' {
		left++
	}
	if left == right {
		return 0
	}
	if !isDigit(s[left]) && !isSign(s[left]) {
		return 0
	}

	res, sign := 0, 1
	if isSign(s[left]) {
		if s[left] == '-' {
			sign = -1
		}
		left++
	}

	// math.MaxIni32: 2147483647
	// math.MinInt32: -2147483648
	for ; left < right && isDigit(s[left]); left++ {
		res = res*10 + int(s[left]-'0')
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && res > math.MaxInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}

func myAtoi2(s string) int {
	automaton := initAutomaton()
	n := len(s)
	for i := 0; i < n; i++ {
		automaton.get(s[i])
	}
	return automaton.sign * automaton.res
}

type Automaton struct {
	sign  int
	res   int
	state string
	table map[string][]string
}

func initAutomaton() Automaton {
	return Automaton{
		sign:  1,
		res:   0,
		state: "start",
		table: map[string][]string{
			"start":     {"start", "signed", "in_number", "end"}, // spaces
			"signed":    {"end", "end", "in_number", "end"},      // sign
			"in_number": {"end", "end", "in_number", "end"},      // numbers
			"end":       {"end", "end", "end", "end"},            // other
		},
	}
}

func (a *Automaton) get(c byte) {
	a.state = a.table[a.state][getCol(c)]
	if a.state == "in_number" {
		a.res = a.res*10 + int(c-'0')
		if a.sign == 1 {
			a.res = min(a.res, math.MaxInt32)
		} else {
			a.res = min(a.res, -math.MinInt32)
		}
	} else if a.state == "signed" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func getCol(c byte) int {
	if c == ' ' {
		return 0
	}
	if isSign(c) {
		return 1
	}
	if isDigit(c) {
		return 2
	}
	return 3
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
