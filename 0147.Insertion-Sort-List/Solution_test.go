package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
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
			args{&ListNode{Val: 1}},
			&ListNode{Val: 1},
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5})},
			structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 1, 2, 5, 5, 4, 10, 0})},
			structures.GenerateListNodesByArray([]int{0, 1, 1, 2, 4, 5, 5, 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
