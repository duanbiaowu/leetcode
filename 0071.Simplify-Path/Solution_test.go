package leetcode

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"/.hidden"},
			"/.hidden",
		},
		{
			"test-2",
			args{"/..hidden"},
			"/..hidden",
		},
		{
			"test-3",
			args{"/abc/..."},
			"/abc/...",
		},
		{
			"test-4",
			args{"/home/"},
			"/home",
		},
		{
			"test-5",
			args{"/..."},
			"/...",
		},
		{
			"test-6",
			args{"/home//foo/"},
			"/home/foo",
		},
		{
			"test-7",
			args{"/a/./b/../../c/"},
			"/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
