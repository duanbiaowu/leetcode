package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{}},
			&ListNode{},
		},
		{
			"test-3",
			args{&ListNode{Val: 1}},
			&ListNode{Val: 1},
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5})},
			&ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5, 6})},
			&ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
