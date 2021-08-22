package leetcode

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
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
			args{"1", ""},
			"1",
		},
		{
			"test-3",
			args{"", "1"},
			"1",
		},
		{
			"test-4",
			args{"1", "1"},
			"10",
		},
		{
			"test-5",
			args{"11", "1"},
			"100",
		},
		{
			"test-6",
			args{"1010", "1011"},
			"10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
