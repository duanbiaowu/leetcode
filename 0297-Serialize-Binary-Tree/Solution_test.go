package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func TestCodec_serialize2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{nil},
			"null,",
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			"1,null,null,",
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			"1,2,null,null,3,null,null,",
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, structures.NULL, 5, structures.NULL, 7})},
			"1,2,null,5,null,null,3,null,7,null,null,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize2(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize2(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{"null,"},
			nil,
		},
		{
			"test-2",
			args{"1,null,null"},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{"1,2,null,null,3,null,null,"},
			structures.GenerateTreeNodesBySlice([]int{1, 2, 3}),
		},
		{
			"test-4",
			args{"1,2,null,5,null,null,3,null,7,null,null,"},
			structures.GenerateTreeNodesBySlice([]int{1, 2, 3, structures.NULL, 5, structures.NULL, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize2(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{nil},
			"[]",
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			"[1,null,null]",
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			"[1,2,3,null,null,null,null]",
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, structures.NULL, 5, structures.NULL, 7})},
			"[1,2,3,null,5,null,7,null,null,null,null]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{""},
			nil,
		},
		{
			"test-2",
			args{"[]"},
			nil,
		},
		{
			"test-3",
			args{"[1,null,null]"},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{"[1,2,3,null,null,null,null]"},
			structures.GenerateTreeNodesBySlice([]int{1, 2, 3}),
		},
		{
			"test-4",
			args{"[1,2,3,null,5,null,7,null,null,null,null]"},
			structures.GenerateTreeNodesBySlice([]int{1, 2, 3, structures.NULL, 5, structures.NULL, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
