package a

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		src  []int
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Divide a slice of length 10 into 5 pieces of length 2",
			args: args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 2},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
		},
		{
			name: "Divide a slice of length 10 into three pieces of length 3 and one of length 1",
			args: args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 3},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
		},
		{
			name: "a slice of length 2 is too short to chunk size 5",
			args: args{src: []int{1, 2}, size: 5},
			want: [][]int{{1, 2}},
		},
		{
			name: "nil1",
			args: args{src: nil, size: 3},
			want: nil,
		},
		{
			name: "nil2",
			args: args{src: []int{}, size: 3},
			want: [][]int{{}},
		},
		{
			name: "nil3",
			args: args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 0},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.src, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
