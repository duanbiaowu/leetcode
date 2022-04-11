package leetcode

import "testing"

func Test_addStrings(t *testing.T) {
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
			args{"", ""},
			"",
		},
		{
			"test-2",
			args{"0", ""},
			"0",
		},
		{
			"test-3",
			args{"", "0"},
			"0",
		},
		{
			"test-4",
			args{"11", "123"},
			"134",
		},
		{
			"test-5",
			args{"456", "77"},
			"533",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
