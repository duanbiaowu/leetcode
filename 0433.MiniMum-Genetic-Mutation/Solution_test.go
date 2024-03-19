package leetcode

import (
	"testing"
)

func Test_minMutation(t *testing.T) {
	type args struct {
		startGene string
		endGene   string
		bank      []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}},
			1,
		},
		{
			"test-2",
			args{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.startGene, tt.args.endGene, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
