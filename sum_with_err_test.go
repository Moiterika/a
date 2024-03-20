package a

import (
	"reflect"
	"testing"
)

func TestSumWithErr(t *testing.T) {
	type args[T any, S number] struct {
		src []T
		s   func(T) S
	}
	type s struct {
		id  int
		qty float64
	}
	tests := []struct {
		name    string
		args    args[s, float64]
		want    float64
		wantErr bool
	}{
		{
			name: "ok",
			args: args[s, float64]{src: []s{{id: 1, qty: 100.5}, {id: 2, qty: 200.5}, {id: 3, qty: 300.2}}, s: func(e s) float64 {
				return e.qty
			}},
			want:    601.2,
			wantErr: false,
		},
		{
			name: "nil",
			args: args[s, float64]{src: nil, s: func(e s) float64 {
				return e.qty
			}},
			want:    0,
			wantErr: false,
		},
		{
			name: "nil2",
			args: args[s, float64]{src: []s{}, s: func(e s) float64 {
				return e.qty
			}},
			want:    0,
			wantErr: false,
		},
		{
			name:    "ng",
			args:    args[s, float64]{src: []s{{id: 1, qty: 100.5}, {id: 2, qty: 200.5}, {id: 3, qty: 300.2}}, s: nil},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumWithErr(tt.args.src, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
