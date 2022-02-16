package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_convertBiNode(t *testing.T) {
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
			args{structures.GenerateTreeNodesBySlice([]int{4, 2, 5, 1, 3, structures.NULL, 6, 0})},
			structures.GenerateTreeNodesBySlice([]int{0, structures.NULL, 1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBiNode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBiNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
