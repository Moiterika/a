package a

import (
	"cmp"
	"reflect"
	"testing"
)

func TestMinBy(t *testing.T) {
	type args[T any, S cmp.Ordered] struct {
		src []T
		m   func(T) S
	}
	type s struct {
		id  int
		qty float64
	}
	tests := []struct {
		name    string
		args    args[s, float64]
		want    s
		wantErr bool
	}{
		{
			name: "ok",
			args: args[s, float64]{src: []s{{1, 100}, {2, 50}, {3, 1000}, {4, -999.5}, {5, 0}}, m: func(e s) float64 {
				return e.qty
			}},
			want:    s{4, -999.5},
			wantErr: false,
		},
		{
			name: "nil1",
			args: args[s, float64]{src: []s{}, m: func(e s) float64 {
				return e.qty
			}},
			want:    s{},
			wantErr: false,
		},
		{
			name: "nil2",
			args: args[s, float64]{src: nil, m: func(e s) float64 {
				return e.qty
			}},
			want:    s{},
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args[s, float64]{src: []s{{1, 100}, {2, 50}, {3, 1000}, {4, -999.5}, {5, 0}}, m: nil},
			want:    s{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinBy(tt.args.src, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
