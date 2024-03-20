package a

import (
	"cmp"
	"reflect"
	"testing"
)

func TestOrderByDescendingWithErr(t *testing.T) {
	type args[T any, S cmp.Ordered] struct {
		src []T
		o   func(T) S
	}
	tests := []struct {
		name    string
		args    args[int, int]
		want    []int
		wantErr bool
	}{
		{
			name: "int-ok",
			args: args[int, int]{src: []int{1, 2, 3, 4, 5}, o: func(e int) int {
				return e
			}},
			want:    []int{5, 4, 3, 2, 1},
			wantErr: false,
		},
		{
			name: "int-nil1",
			args: args[int, int]{src: nil, o: func(e int) int {
				return e
			}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "int-nil2",
			args: args[int, int]{src: []int{}, o: func(e int) int {
				return e
			}},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args[int, int]{src: []int{1, 2, 3, 4, 5}, o: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrderByDescendingWithErr(tt.args.src, tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderByDescending() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderByDescending() = %v, want %v", got, tt.want)
			}
		})
	}
}
