package a

import "testing"

func TestAny(t *testing.T) {
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
				return e >= 5
			}},
			want: true,
		},
		{
			name: "ok2",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a: func(e int) bool {
					return e < 0
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
			want: false,
		},
		{
			name: "nil2",
			args: args{
				src: nil,
				a: func(e int) bool {
					return e <= 4
				},
			},
			want: false,
		},
		{
			name: "ng-nil3;irregular",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a:   nil, // irregular
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.args.src, tt.args.a)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}
