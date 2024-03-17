package a

import (
	"reflect"
	"testing"
)

func TestLastOrDefault(t *testing.T) {
	type args[T any] struct {
		src []T
	}
	src := []int{300, 200, 100}

	tests := []struct {
		name string
		args args[*int]
		want *int
	}{
		{
			name: "ok",
			args: args[*int]{src: []*int{&src[0], &src[1], &src[2]}},
			want: &src[2],
		},
		{
			name: "ok-nil",
			args: args[*int]{src: []*int{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastOrDefault(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
