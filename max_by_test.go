package a

import (
	"cmp"
	"reflect"
	"testing"
)

func TestMaxBy(t *testing.T) {
	type args[T any, S cmp.Ordered] struct {
		src []T
		m   func(T) S
	}
	type s struct {
		id  int
		qty float64
	}
	tests := []struct {
		name string
		args args[s, float64]
		want s
	}{
		{
			name: "ok",
			args: args[s, float64]{
				src: []s{
					{1, 100},
					{2, 50},
					{3, 1000},
					{4, -1000.5},
					{5, 0},
				},
				m: func(e s) float64 {
					return e.qty
				},
			},
			want: s{3, 1000},
		},
		{
			name: "nil1",
			args: args[s, float64]{
				src: []s{},
				m: func(e s) float64 {
					return e.qty
				},
			},
			want: s{},
		},
		{
			name: "nil2",
			args: args[s, float64]{
				src: nil,
				m: func(e s) float64 {
					return e.qty
				},
			},
			want: s{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxBy(tt.args.src, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
