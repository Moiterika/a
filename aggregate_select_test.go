package a

import (
	"math"
	"reflect"
	"testing"
)

func TestAggregateSelect(t *testing.T) {
	type args[T, S any] struct {
		src      []T
		initialF func(T) S
		f        func(S, T) S
	}
	intTests := []struct {
		name string
		args[int, int]
		want int
	}{
		{
			name: "ok-sum",
			args: args[int, int]{
				src: []int{1, 2, 3},
				initialF: func(e int) int {
					return e
				},
				f: func(rlt int, cur int) int {
					return rlt + cur
				},
			},
			want: 6,
		},
		{
			name: "ok-double-sum",
			args: args[int, int]{
				src: []int{1, 2, 3},
				initialF: func(e int) int {
					return 2 * e
				},
				f: func(rlt int, cur int) int {
					return rlt + 2*cur
				},
			},
			want: 12,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AggregateSelect(tt.args.src, tt.args.initialF, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSelect() = %v, want %v", got, tt.want)
			}
		})
	}

	type result struct {
		sum float64
		max float64
		min float64
	}
	structTests := []struct {
		name string
		args[int, result]
		want result
	}{
		{
			name: "ok-struct",
			args: args[int, result]{
				src: []int{1, 2, 3},
				initialF: func(e int) result {
					return result{
						sum: float64(e),
						max: float64(e),
						min: float64(e),
					}
				},
				f: func(rlt result, cur int) result {
					c := float64(cur)
					rlt.sum += c
					rlt.max = math.Max(rlt.max, c)
					rlt.min = math.Min(rlt.min, c)
					return rlt
				},
			},
			want: result{
				sum: 6,
				max: 3,
				min: 1,
			},
		},
		{
			name: "ok-original-initial-value",
			args: args[int, result]{
				src: []int{1, 2, 3},
				initialF: func(_ int) result {
					return result{
						sum: 10,
						max: 99,
						min: 0,
					}
				},
				f: func(rlt result, cur int) result {
					c := float64(cur)
					rlt.sum += c
					rlt.max = math.Max(rlt.max, c)
					rlt.min = math.Min(rlt.min, c)
					return rlt
				},
			},
			want: result{
				sum: 15,
				max: 99,
				min: 0,
			},
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AggregateSelect(tt.args.src, tt.args.initialF, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
