package leetcode

import (
	"reflect"
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{},
			nil,
		},
		{
			"test-2",
			args{&structures.TreeNode{Val: 0}},
			[]int{0},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, structures.NULL, 2, 2})},
			[]int{2},
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, structures.NULL, 2})},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
