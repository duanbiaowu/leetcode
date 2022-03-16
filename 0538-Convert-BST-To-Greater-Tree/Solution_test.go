package leetcode

import (
	"reflect"
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_convertBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{4, 1, 6, 0, 2, 5, 7, structures.NULL, structures.NULL, structures.NULL, 3, structures.NULL, structures.NULL, structures.NULL, 8})},
			structures.GenerateTreeNodesBySlice([]int{30, 36, 21, 36, 35, 26, 15, structures.NULL, structures.NULL, structures.NULL, 33, structures.NULL, structures.NULL, structures.NULL, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
