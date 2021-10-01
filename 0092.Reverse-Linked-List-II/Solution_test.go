package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, 1, 1},
			nil,
		},
		{
			"test-2",
			args{&ListNode{}, 1, 1},
			&ListNode{},
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{3, 5}), 1, 2},
			structures.GenerateListNodesByArray([]int{5, 3}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 2, 4},
			structures.GenerateListNodesByArray([]int{1, 4, 3, 2, 5}),
		},
		{
			"test-5",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 1, 5},
			structures.GenerateListNodesByArray([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
