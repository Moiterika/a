package a

import (
	"cmp"
	"reflect"
	"testing"
)

func TestOrderBy(t *testing.T) {
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
			args: args[int, int]{src: []int{5, 4, 3, 2, 1}, o: func(e int) int {
				return e
			}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}

	type s struct {
		x    int
		note string
	}
	structTests := []struct {
		name string
		args args[s, int]
		want []s
	}{
		{
			name: "struct-ok",
			args: args[s, int]{src: []s{
				{x: 1, note: "def"},
				{x: 5, note: "zzz"},
				{x: 1, note: "zzz"},
				{x: 1, note: "abc"},
				{x: 2, note: "opq"},
				{x: 5, note: "aaa"},
			}, o: func(e s) int {
				return e.x
			}},
			want: []s{
				{x: 1, note: "def"},
				{x: 1, note: "zzz"},
				{x: 1, note: "abc"},
				{x: 2, note: "opq"},
				{x: 5, note: "zzz"},
				{x: 5, note: "aaa"},
			},
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}

	structTests2 := []struct {
		name string
		args args[s, string]
		want []s
	}{
		{
			name: "struct-ok2",
			args: args[s, string]{src: []s{
				{x: 1, note: "def"},
				{x: 5, note: "zzz"},
				{x: 1, note: "zzz"},
				{x: 1, note: "abc"},
				{x: 2, note: "opq"},
				{x: 5, note: "aaa"},
			}, o: func(e s) string {
				return e.note
			}},
			want: []s{
				{x: 5, note: "aaa"},
				{x: 1, note: "abc"},
				{x: 1, note: "def"},
				{x: 2, note: "opq"},
				{x: 5, note: "zzz"},
				{x: 1, note: "zzz"},
			},
		},
	}
	for _, tt := range structTests2 {
		t.Run(tt.name, func(t *testing.T) {
			OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(tt.args.src, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", tt.args.src, tt.want)
			}
		})
	}
}
