package a

import (
	"reflect"
	"testing"
)

func TestDistinctBy(t *testing.T) {
	type args[T any, K comparable] struct {
		src []T
		d   func(T) K
	}
	type s struct {
		id   int
		note string
	}
	tests := []struct {
		name string
		args args[s, int]
		want []s
	}{
		{
			name: "int",
			args: args[s, int]{src: []s{{1, "z"}, {2, "b"}, {1, "a"}, {3, "c"}, {2, "a"}}, d: func(e s) int {
				return e.id
			}},
			want: []s{{1, "z"}, {2, "b"}, {3, "c"}},
		},
		{
			name: "int-nil1",
			args: args[s, int]{src: nil, d: func(e s) int {
				return e.id
			}},
			want: nil,
		},
		{
			name: "int-nil2",
			args: args[s, int]{src: []s{}, d: func(e s) int {
				return e.id
			}},
			want: []s{},
		},
		{
			name: "int-func-nil",
			args: args[s, int]{src: []s{}, d: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.src, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
