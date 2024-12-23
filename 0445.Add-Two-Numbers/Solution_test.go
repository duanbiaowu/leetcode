package leetcode

import (
	"reflect"
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{Val: 1}, nil},
			&ListNode{Val: 1},
		},
		{
			"test-3",
			args{nil, &ListNode{Val: 2}},
			&ListNode{Val: 2},
		},
		{
			"test-4",
			args{&ListNode{Val: 1}, &ListNode{Val: 2}},
			&ListNode{Val: 3},
		},
		{
			"test-5",
			args{
				structures.GenerateListNodesByArray([]int{7, 2, 4, 3}),
				structures.GenerateListNodesByArray([]int{5, 6, 4}),
			},
			structures.GenerateListNodesByArray([]int{7, 8, 0, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
