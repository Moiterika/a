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
		name       string
		args       args[int]
		wantRet    [][]int
		wantOthers []int
	}{
		{
			"int-ok",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return e >= 3 }},
			},
			[][]int{{3, 4, 5}},
			[]int{1, 2},
		},
		{
			"int-nil",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return false }},
			},
			[][]int{nil},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"int-overlapping filters",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				[]func(e int) bool{func(e int) bool { return e >= 3 }, func(e int) bool { return e >= 4 }},
			},
			[][]int{{3, 4, 5}, {4, 5}},
			[]int{1, 2},
		},
		{
			"int-no filters",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				nil,
			},
			nil,
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, gotOthers := Filters(tt.args.src, tt.args.fs...)
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("Filters()'s Filtered Results = %v, want %v", gotRet, tt.wantRet)
			}
			if !reflect.DeepEqual(gotOthers, tt.wantOthers) {
				t.Errorf("Filters()'s Unfiltered Result = %v, want %v", gotOthers, tt.wantOthers)
			}
		})
	}
}
