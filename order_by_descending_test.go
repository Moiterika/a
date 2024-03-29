package a

import (
	"cmp"
	"reflect"
	"testing"
)

func TestOrderByDescending(t *testing.T) {
	type args[T any, S cmp.Ordered] struct {
		src []T
		o   func(T) S
	}
	tests := []struct {
		name string
		args args[int, int]
		want []int
	}{
		{
			name: "int-ok",
			args: args[int, int]{src: []int{1, 2, 3, 4, 5}, o: func(e int) int {
				return e
			}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "int-nil1",
			args: args[int, int]{src: nil, o: func(e int) int {
				return e
			}},
			want: nil,
		},
		{
			name: "int-nil2",
			args: args[int, int]{src: []int{}, o: func(e int) int {
				return e
			}},
			want: []int{},
		},
		{
			name: "int-func-nil",
			args: args[int, int]{src: []int{1, 2, 3, 4, 5}, o: nil},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderByDescending(tt.args.src, tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
			}
		})
	}
}
