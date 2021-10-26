package leetcode

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"0", "0"},
			"0",
		},
		{
			"test-2",
			args{"2", "3"},
			"6",
		},
		{
			"test-3",
			args{"123", "456"},
			"56088",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
