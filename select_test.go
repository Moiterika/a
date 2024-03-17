package a

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	type args[T, S any] struct {
		src []T
		s   func(T) S
	}
	tests := []struct {
		name string
		args args[int, int]
		want []int
	}{
		{
			name: "ok",
			args: args[int, int]{src: []int{1, 2, 3}, s: func(e int) int { return e * 2 }},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.src, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
