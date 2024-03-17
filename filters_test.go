package a

import (
	"reflect"
	"testing"
)

func TestFilters(t *testing.T) {
	type args[T any] struct {
		src []T
		fs  []func(T) bool
	}
	tests := []struct {
		name string
		args args[int]
		want [][]int
	}{
		{
			"int-ok",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return e >= 3 }},
			},
			[][]int{{3, 4, 5}, {1, 2}},
		},
		{
			"int-nil",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return false }},
			},
			[][]int{nil, {1, 2, 3, 4, 5}},
		},
		{
			"int-overlapping filters",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return e >= 3 }, func(e int) bool { return e >= 4 }},
			},
			[][]int{{3, 4, 5}, {4, 5}, {1, 2}},
		},
		{
			"int-no filters",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				nil,
			},
			[][]int{{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filters(tt.args.src, tt.args.fs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filters() = %v, want %v", got, tt.want)
			}
		})
	}
}
