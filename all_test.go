package a

import "testing"

func TestAll(t *testing.T) {
	type args struct {
		src []int
		a   func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{src: []int{1, 2, 3, 4, 5}, a: func(e int) bool {
				return e > 0
			}},
			want: true,
		},
		{
			name: "ok2",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a: func(e int) bool {
					return e <= 4
				},
			},
			want: false,
		},
		{
			name: "nil1",
			args: args{
				src: []int{},
				a: func(e int) bool {
					return e <= 4
				},
			},
			want: true,
		},
		{
			name: "nil2",
			args: args{
				src: nil,
				a: func(e int) bool {
					return e <= 4
				},
			},
			want: true,
		},
		{
			name: "ng-nil;irregular",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a:   nil, // irregular
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.args.src, tt.args.a)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
