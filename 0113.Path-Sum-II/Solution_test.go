package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil, 0},
			nil,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}, 1},
			[][]int{
				{1},
			},
		},
		{
			"test-3",
			args{
				structures.GenerateTreeNodesBySlice([]int{1, 2, 3}),
				5,
			},
			nil,
		},
		{
			"test-4",
			args{
				structures.GenerateTreeNodesBySlice([]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1}),
				22,
			},
			[][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
