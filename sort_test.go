package a

import (
	"reflect"
	"testing"
)

type s struct {
	x    int
	note string
}

func TestSort(t *testing.T) {
	type args[T any] struct {
		src  []T
		cmps []func(t1 T, t2 T) bool
	}
	tests := []struct {
		name string
		args args[s]
		want []s
	}{
		{
			name: "single-sort-ordered",
			args: args[s]{
				src: []s{
					{x: 1, note: "def"},
					{x: 3, note: "abc"},
					{x: 2, note: "def"},
					{x: 2, note: "abc"},
				},
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.x < t2.x
					},
				},
			},
			want: []s{
				{x: 1, note: "def"},
				{x: 2, note: "def"},
				{x: 2, note: "abc"},
				{x: 3, note: "abc"},
			},
		},
		{
			name: "single-sort-ordered 2",
			args: args[s]{
				src: []s{
					{x: 1, note: "def"},
					{x: 3, note: "abc"},
					{x: 2, note: "def"},
					{x: 2, note: "abc"},
				},
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.note < t2.note
					},
				},
			},
			want: []s{
				{x: 3, note: "abc"},
				{x: 2, note: "abc"},
				{x: 1, note: "def"},
				{x: 2, note: "def"},
			},
		},
		{
			name: "single-sort-descending-ordered",
			args: args[s]{
				src: []s{
					{x: 1, note: "def"},
					{x: 3, note: "abc"},
					{x: 2, note: "def"},
					{x: 2, note: "abc"},
				},
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.x > t2.x // 逆順
					},
				},
			},
			want: []s{
				{x: 3, note: "abc"},
				{x: 2, note: "def"},
				{x: 2, note: "abc"},
				{x: 1, note: "def"},
			},
		},
		{
			name: "multi-sort",
			args: args[s]{
				src: []s{
					{x: 1, note: "def"},
					{x: 3, note: "abc"},
					{x: 2, note: "def"},
					{x: 2, note: "abc"},
				},
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.x < t2.x
					},
					func(t1 s, t2 s) bool {
						return t1.note < t2.note
					},
				},
			},
			want: []s{
				{x: 1, note: "def"},
				{x: 2, note: "abc"},
				{x: 2, note: "def"},
				{x: 3, note: "abc"},
			},
		},
		{
			name: "src is nil",
			args: args[s]{
				src: nil,
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.x < t2.x
					},
				},
			},
			want: nil,
		},
		{
			name: "len(src) == 1",
			args: args[s]{
				src: []s{
					{x: 1, note: "def"},
				},
				cmps: []func(t1 s, t2 s) bool{
					func(t1 s, t2 s) bool {
						return t1.x < t2.x
					},
				},
			},
			want: []s{
				{x: 1, note: "def"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.src, tt.args.cmps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
