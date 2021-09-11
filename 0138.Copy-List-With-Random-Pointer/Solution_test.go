package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&Node{Val: 1}},
			&Node{Val: 1},
		},
		{
			"test-3",
			args{structures.GenerateRandomListNodesBySlice([][2]int{
				{1, 0},
				{2, 0},
				{3, 0},
			})},
			structures.GenerateRandomListNodesBySlice([][2]int{
				{1, 0},
				{2, 0},
				{3, 0},
			}),
		},
		{
			"test-4",
			args{structures.GenerateRandomListNodesBySlice([][2]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{4, 2},
				{5, 0},
			})},
			structures.GenerateRandomListNodesBySlice([][2]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{4, 2},
				{5, 0},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
