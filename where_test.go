package a

import (
	"reflect"
	"testing"
)

func TestWhere(t *testing.T) {
	// int型OKテスト
	type args[T any] struct {
		src []T
		w   func(T) bool
	}
	intTests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			"int-ok",
			args[int]{
				[]int{1, 2, 3, 4, 5},
				func(e int) bool { return e >= 3 },
			},
			[]int{3, 4, 5},
		},
		{
			"int-nil1",
			args[int]{
				nil,
				func(e int) bool { return e >= 3 },
			},
			[]int{}, // nil => error because typed-nil
		},
		{
			"int-nil2",
			args[int]{
				[]int{},
				func(e int) bool { return e >= 3 },
			},
			[]int{},
		},
		{
			"int-nil3",
			args[int]{
				nil,
				func(e int) bool { return e >= 3 },
			},
			[]int{},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Where(tt.args.src, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}

	// *int型OKテスト
	src := []int{1, 2, 3, 4, 5}
	intPtrTests := []struct {
		name string
		args args[*int]
		want []*int
	}{
		{
			"*int-ok",
			args[*int]{
				[]*int{&src[0], &src[1], &src[2], &src[3], &src[4]},
				func(e *int) bool { return *e >= 3 },
			},
			[]*int{&src[2], &src[3], &src[4]},
		},
		{
			"*int-nil",
			args[*int]{
				[]*int{&src[0], &src[1], &src[2], &src[3], &src[4]},
				func(e *int) bool { return *e < 0 },
			},
			[]*int{}, // nil => error because typed-nil
		},
	}
	for _, tt := range intPtrTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Where(tt.args.src, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}
