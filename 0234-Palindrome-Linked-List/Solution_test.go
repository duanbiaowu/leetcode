package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil},
			true,
		},
		{
			"test-2",
			args{&ListNode{}},
			true,
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2})},
			false,
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 2, 1})},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
