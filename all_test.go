package a

import "testing"

func TestAll(t *testing.T) {
	type args struct {
		src []int
		a   func(int) bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ok",
			args: args{src: []int{1, 2, 3, 4, 5}, a: func(e int) bool {
				return e > 0
			}},
			want:    true,
			wantErr: false,
		},
		{
			name: "ok2",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a: func(e int) bool {
					return e <= 4
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "nil1",
			args: args{
				src: []int{},
				a: func(e int) bool {
					return e <= 4
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "nil2",
			args: args{
				src: nil,
				a: func(e int) bool {
					return e <= 4
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "nil3",
			args: args{
				src: []int{1, 2, 3, 4, 5},
				a:   nil,
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := All(tt.args.src, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
