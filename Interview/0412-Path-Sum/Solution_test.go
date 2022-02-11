package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, 0},
			0,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, 5, 1}), 22},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
