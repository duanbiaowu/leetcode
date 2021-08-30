package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{1},
			"A",
		},
		{
			"test-2",
			args{28},
			"AB",
		},
		{
			"test-3",
			args{701},
			"ZY",
		},
		{
			"test-4",
			args{2147483647},
			"FXSHRXW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
