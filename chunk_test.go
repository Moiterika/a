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
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name:    "Divide a slice of length 10 into 5 pieces of length 2",
			args:    args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 2},
			want:    [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			wantErr: false,
		},
		{
			name:    "Divide a slice of length 10 into three pieces of length 3 and one of length 1",
			args:    args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 3},
			want:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
			wantErr: false,
		},
		{
			name:    "a slice of length 2 is too short to chunk size 5",
			args:    args{src: []int{1, 2}, size: 5},
			want:    [][]int{{1, 2}},
			wantErr: false,
		},
		{
			name:    "nil1",
			args:    args{src: nil, size: 3},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "nil2",
			args:    args{src: []int{}, size: 3},
			want:    [][]int{{}},
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{src: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, size: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Chunk(tt.args.src, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
