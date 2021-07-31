package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		first *ListNode
		last  *ListNode
	}

	head := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	fourth := &ListNode{Val: 4}

	head.Next = second
	head.Next.Next = third
	head.Next.Next.Next = fourth

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{}, nil},
			&ListNode{},
		},
		{
			"test-3",
			args{head, fourth},
			third,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.first, tt.args.last); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, 0},
			nil,
		},
		{
			"test-2",
			args{nil, 1},
			nil,
		},
		{
			"test-3",
			args{nil, -1},
			nil,
		},
		{
			"test-4",
			args{&ListNode{}, 0},
			&ListNode{},
		},
		{
			"test-5",
			args{&ListNode{Val: 1}, 1},
			&ListNode{Val: 1},
		},
		{
			"test-6",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3}), 1},
			structures.GenerateListNodesByArray([]int{1, 2, 3}),
		},
		{
			"test-7",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 2},
			structures.GenerateListNodesByArray([]int{2, 1, 4, 3, 5}),
		},
		{
			"test-8",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5}), 3},
			structures.GenerateListNodesByArray([]int{3, 2, 1, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
