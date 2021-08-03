package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, 0},
			nil,
		},
		{
			"test-2",
			args{&ListNode{Val: 0}, 1},
			&ListNode{Val: 0},
		},
		{
			"test-3",
			args{&ListNode{Val: 1}, 0},
			&ListNode{Val: 1},
		},
		{
			"test-4",
			args{&ListNode{Val: 1}, 1},
			nil,
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3}), 1},
			structures.GenerateListNodesByArray([]int{2, 3}),
		},
		{
			"test-6",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3}), 2},
			structures.GenerateListNodesByArray([]int{1, 3}),
		},
		{
			"test-7",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3}), 3},
			structures.GenerateListNodesByArray([]int{1, 2}),
		},
		{
			"test-8",
			args{structures.GenerateListNodesByArray([]int{1, 1, 3}), 1},
			structures.GenerateListNodesByArray([]int{3}),
		},
		{
			"test-9",
			args{structures.GenerateListNodesByArray([]int{1, 1, 2, 3, 4, 1, 5}), 1},
			structures.GenerateListNodesByArray([]int{2, 3, 4, 5}),
		},
		{
			"test-10",
			args{structures.GenerateListNodesByArray([]int{1, 1, 1, 1, 1}), 1},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
