package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"test-1",
			args{nil},
			[][]string{},
		},
		{
			"test-2",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) == len(got[j]) {
					sort.Strings(got[i])
					sort.Strings(got[j])

					for k := range got[i] {
						if len(got[i][j]) > len(got[j][k]) {
							return true
						}
					}

					return false
				}
				return len(got[i]) > len(got[j])
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
