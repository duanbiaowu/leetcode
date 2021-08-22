package leetcode

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{[]int{}},
			[][]int{
				{},
			},
		},
		{
			"test-2",
			args{[]int{1}},
			[][]int{
				{},
				{1},
			},
		},
		{
			"test-3",
			args{[]int{1, 2}},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsetsRecursively(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{[]int{}},
			[][]int{
				{},
			},
		},
		{
			"test-2",
			args{[]int{1}},
			[][]int{
				{},
				{1},
			},
		},
		{
			"test-3",
			args{[]int{1, 2}},
			[][]int{
				{},
				{1},
				{1, 2},
				{2},
			},
		},
		{
			"test-4",
			args{[]int{1, 2, 3}},
			[][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsRecursively(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}