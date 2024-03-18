package a

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	type args[T comparable] struct {
		src []T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "int",
			args: args[int]{src: []int{1, 2, 1, 5, 2, 3, 1, 4}},
			want: []int{1, 2, 5, 3, 4},
		},
		{
			name: "int-nil1",
			args: args[int]{src: nil},
			want: nil,
		},
		{
			name: "int-nil2",
			args: args[int]{src: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
