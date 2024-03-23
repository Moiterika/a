package a

import (
	"reflect"
	"testing"
)

func TestAggregateSelectWithIndex(t *testing.T) {
	type s struct {
		i   int
		val float64
	}
	type args[T, S any] struct {
		src      []T
		initialF func(T) S
		f        func(S, T, int) S
	}
	tests := []struct {
		name string
		args args[float64, s]
		want s
	}{
		{
			name: "",
			args: args[float64, s]{
				src: []float64{1, 2, 3, 1, 3},
				initialF: func(e float64) s {
					return s{
						i:   0,
						val: e,
					}
				},
				f: func(rlt s, cur float64, idx int) s {
					// index and value of max in src
					if rlt.val < cur {
						rlt.i = idx
						rlt.val = cur
					}
					return rlt
				},
			},
			want: s{
				i:   2,
				val: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AggregateSelectWithIndex(tt.args.src, tt.args.initialF, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSelectWithIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
