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
	intTests := []struct {
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
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			OrderByDescending(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("OrderByDescending() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}
}
