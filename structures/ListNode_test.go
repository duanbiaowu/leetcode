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
