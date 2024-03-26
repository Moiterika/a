package a

import (
	"reflect"
	"testing"
)

func TestSelectMany(t *testing.T) {
	type args[T, S any] struct {
		src []T
		s   func(T) []S
	}
	tests := []struct {
		name string
		args args[[]int, int]
		want []int
	}{
		{
			name: "ok",
			args: args[[]int, int]{
				src: [][]int{{1, 2, 3}, {3, 4}, {5}},
				s: func(e []int) []int {
					return e
				},
			},
			want: []int{1, 2, 3, 3, 4, 5},
		},
		{
			name: "ok-nil",
			args: args[[]int, int]{
				src: [][]int{{1, 2, 3}, nil, {}, {5}},
				s: func(e []int) []int {
					return e
				},
			},
			want: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectMany(tt.args.src, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectMany() = %v, want %v", got, tt.want)
			}
		})
	}
}
