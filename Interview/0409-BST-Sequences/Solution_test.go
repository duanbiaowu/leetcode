package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func TestBSTSequences(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil},
			[][]int{},
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{1})},
			[][]int{{1}},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{2, 1, 3})},
			[][]int{
				{2, 1, 3},
				{2, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BSTSequences(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
