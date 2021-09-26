package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	//root := structures.GenerateTreeNodesBySlice([]int{
	//	1, structures.NULL, 2, structures.NULL, 5,
	//})
	//
	//fmt.Println( root.Right.Right )
	//return

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
			args{&TreeNode{Val: 0}},
			&TreeNode{Val: 0},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{
				1, 2, 5, 3, 4, structures.NULL, 6,
			})},
			structures.GenerateTreeNodesBySlice([]int{
				1,
				structures.NULL, 2,
				structures.NULL, 3,
				structures.NULL, 4,
				structures.NULL, 5,
				structures.NULL, 6,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("connect() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
