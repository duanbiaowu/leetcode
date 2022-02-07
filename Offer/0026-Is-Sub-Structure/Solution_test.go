package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil, nil},
			false,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3}), nil},
			false,
		},
		{
			"test-3",
			args{nil, structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			false,
		},
		{
			"test-4",
			args{
				structures.GenerateTreeNodesBySlice([]int{1, 2, 3}),
				structures.GenerateTreeNodesBySlice([]int{3, 1}),
			},
			false,
		},
		{
			"test-5",
			args{
				structures.GenerateTreeNodesBySlice([]int{3, 4, 5, 1, 2}),
				structures.GenerateTreeNodesBySlice([]int{4, 1}),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
