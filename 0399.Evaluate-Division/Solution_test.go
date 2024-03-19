package leetcode

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			/**

			输入： equations = [["a","b"],["b","c"]]
			      values = [2.0,3.0]
				  queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]

			输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]

			解释：
			条件：a / b = 2.0, b / c = 3.0
			问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
			结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
			注意：x 是未定义的 => -1.0

			*/
			"test-1",
			args{
				equations: [][]string{
					{"a", "b"},
					{"b", "c"},
				},
				values: []float64{2, 3},
				queries: [][]string{
					{"a", "c"},
					{"b", "a"},
					{"a", "e"},
					{"a", "a"},
					{"x", "x"},
				},
			},
			[]float64{6, 0.5, -1, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
