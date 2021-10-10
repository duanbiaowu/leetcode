package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
			args{&ListNode{}, 1},
			&ListNode{},
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 2},
			structures.GenerateListNodesByArray([]int{4, 5, 1, 2, 3}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 3},
			structures.GenerateListNodesByArray([]int{3, 4, 5, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
