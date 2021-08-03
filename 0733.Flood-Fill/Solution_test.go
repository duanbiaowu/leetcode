package leetcode

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{
				[][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				1,
				1,
				2,
			},
			[][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
