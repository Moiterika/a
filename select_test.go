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
		{
			name: "nil1",
			args: args[int, int]{src: nil, s: func(e int) int { return e * 2 }},
			want: nil,
		},
		{
			name: "nil2",
			args: args[int, int]{src: []int{}, s: func(e int) int { return e * 2 }},
			want: []int{},
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
