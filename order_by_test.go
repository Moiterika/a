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
		{
			name: "int-nil",
			args: args[int, int]{src: []int{}, o: func(e int) int {
				return e
			}},
			want: []int{},
		},
		{
			name: "int-nil2",
			args: args[int, int]{src: nil, o: func(e int) int {
				return e
			}},
			want: nil,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}

	src := []int{3, 4, 5, 1, 2}
	intPtrTests := []struct {
		name string
		args args[*int, int]
		want []*int
	}{
		{
			name: "*int-ok",
			args: args[*int, int]{src: []*int{&src[0], &src[1], &src[2], &src[3], &src[4]}, o: func(e *int) int {
				return *e
			}},
			want: []*int{&src[3], &src[4], &src[0], &src[1], &src[2]},
		},
	}
	for _, tt := range intPtrTests {
		t.Run(tt.name, func(t *testing.T) {
			got := OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
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
			got := OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
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
			got := OrderBy(tt.args.src, tt.args.o)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
