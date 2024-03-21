package a

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type s struct {
		x    int
		note string
	}
	type args[T any, K comparable] struct {
		src []T
		k   func(T) K
	}
	groupingIntTests := []struct {
		name string
		args args[s, int]
		want map[int][]s
	}{
		{
			name: "ok-int",
			args: args[s, int]{
				src: []s{
					{x: 1, note: "def"},
					{x: 5, note: "zzz"},
					{x: 1, note: "zzz"},
					{x: 1, note: "abc"},
					{x: 2, note: "opq"},
					{x: 5, note: "aaa"},
				},
				k: func(e s) int {
					return e.x
				},
			},
			want: map[int][]s{
				1: {
					{x: 1, note: "def"},
					{x: 1, note: "zzz"},
					{x: 1, note: "abc"},
				},
				5: {
					{x: 5, note: "zzz"},
					{x: 5, note: "aaa"},
				},
				2: {
					{x: 2, note: "opq"},
				},
			},
		},
		{
			name: "nil",
			args: args[s, int]{
				src: []s{},
				k: func(e s) int {
					return e.x
				},
			},
			want: map[int][]s{},
		},
		{
			name: "nil2",
			args: args[s, int]{
				src: nil,
				k: func(e s) int {
					return e.x
				},
			},
			want: map[int][]s{},
		},
		{
			name: "nil3",
			args: args[s, int]{
				src: []s{
					{x: 1, note: "def"},
					{x: 5, note: "zzz"},
					{x: 1, note: "zzz"},
					{x: 1, note: "abc"},
					{x: 2, note: "opq"},
					{x: 5, note: "aaa"},
				},
				k: nil,
			},
			want: map[int][]s{},
		},
		{
			name: "nil4",
			args: args[s, int]{
				src: nil,
				k:   nil,
			},
			want: map[int][]s{},
		},
	}
	for _, tt := range groupingIntTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.src, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}

	groupingStringTests := []struct {
		name string
		args args[s, string]
		want map[string][]s
	}{
		{
			name: "ok-string",
			args: args[s, string]{
				src: []s{
					{x: 1, note: "def"},
					{x: 5, note: "zzz"},
					{x: 1, note: "zzz"},
					{x: 1, note: "abc"},
					{x: 2, note: "opq"},
					{x: 5, note: "aaa"},
				},
				k: func(e s) string {
					return e.note
				},
			},
			want: map[string][]s{
				"def": {
					{x: 1, note: "def"},
				},
				"zzz": {
					{x: 5, note: "zzz"},
					{x: 1, note: "zzz"},
				},
				"abc": {
					{x: 1, note: "abc"},
				},
				"opq": {
					{x: 2, note: "opq"},
				},
				"aaa": {
					{x: 5, note: "aaa"},
				},
			},
		},
	}
	for _, tt := range groupingStringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.src, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
