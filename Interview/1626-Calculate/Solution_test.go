package leetcode

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"test-1",
		//	args{""},
		//	0,
		//},
		//{
		//	"test-2",
		//	args{" "},
		//	0,
		//},
		{
			"test-3",
			args{"3+2*2"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
