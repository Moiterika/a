package a

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	type args[T, S any] struct {
		src []T
		s   func(T) S
	}
	tests := []struct {
		name    string
		args    args[int, int]
		want    []int
		wantErr bool
	}{
		{
			name: "ok",
			args: args[int, int]{src: []int{1, 2, 3}, s: func(e int) int {
				return e * 2
			}},
			want:    []int{2, 4, 6},
			wantErr: false,
		},
		{
			name: "nil1",
			args: args[int, int]{src: nil, s: func(e int) int {
				return e * 2
			}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "nil2",
			args: args[int, int]{src: []int{}, s: func(e int) int {
				return e * 2
			}},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args[int, int]{src: []int{1, 2, 3}, s: nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.src, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
