package structures

import (
	"reflect"
	"testing"
)

func TestGenerateListNodesByArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test an empty array",
			args{[]int{}},
			nil,
		},
		{
			"test an array with only one element",
			args{[]int{0}},
			&ListNode{0, nil},
		},
		{
			"test an array with multiple elements",
			args{[]int{0, 1, 2}},
			&ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateListNodesByArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateListNodesByArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomListNodesBySlice(t *testing.T) {
	type args struct {
		nums [][2]int
	}

	header := &RandomListNode{Val: 1}
	node2 := &RandomListNode{Val: 2}
	node3 := &RandomListNode{Val: 3}
	node4 := &RandomListNode{Val: 4}
	node5 := &RandomListNode{Val: 5}

	node4.Next = node5
	node3.Next = node4
	node2.Next = node3
	header.Next = node2

	node3.Random = header
	node4.Random = node2

	tests := []struct {
		name string
		args args
		want *RandomListNode
	}{
		{
			"test an empty slice",
			args{nil},
			nil,
		},
		{
			"test an slice with only one element",
			args{[][2]int{
				{1, 0},
			}},
			&RandomListNode{Val: 1},
		},
		{
			"test an slice of multiple elements with no random elements",
			args{[][2]int{
				{1, 0},
				{2, 0},
				{3, 0},
			}},
			&RandomListNode{
				Val: 1,
				Next: &RandomListNode{
					Val: 2,
					Next: &RandomListNode{
						Val: 3,
					},
				},
			},
		},
		{
			"test an slice of multiple elements with random elements",
			args{[][2]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{4, 2},
				{5, 0},
			}},
			header,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomListNodesBySlice(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomListNodesBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
