package a

import (
	"reflect"
	"testing"
)

func TestGroupByWithErr(t *testing.T) {
	type s struct {
		x    int
		note string
	}
	type args[T any, K comparable] struct {
		src []T
		k   func(T) K
	}
	groupingIntTests := []struct {
		name    string
		args    args[s, int]
		want    map[int][]s
		wantErr bool
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
			wantErr: false,
		},
		{
			name: "nil",
			args: args[s, int]{
				src: []s{},
				k: func(e s) int {
					return e.x
				},
			},
			want:    map[int][]s{},
			wantErr: false,
		},
		{
			name: "nil2",
			args: args[s, int]{
				src: nil,
				k: func(e s) int {
					return e.x
				},
			},
			want:    map[int][]s{},
			wantErr: false,
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
			want:    map[int][]s{},
			wantErr: true,
		},
		{
			name: "nil4",
			args: args[s, int]{
				src: nil,
				k:   nil,
			},
			want:    map[int][]s{},
			wantErr: true,
		},
	}
	for _, tt := range groupingIntTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GroupByWithErr(tt.args.src, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupByWithErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByWithErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
