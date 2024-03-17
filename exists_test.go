package a

import "testing"

func TestExists(t *testing.T) {
	type args[T any] struct {
		src []T
		e   T
	}
	intTests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "int-exists",
			args: args[int]{src: []int{1, 2, 3}, e: 2},
			want: true,
		},
		{
			name: "int-not exists",
			args: args[int]{src: []int{1, 2, 3}, e: 5},
			want: false,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.src, tt.args.e); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}

	src := []int{10, 20, 30, 99}
	outOfSrc := 20
	intPtrTests := []struct {
		name string
		args args[*int]
		want bool
	}{
		{
			name: "*int-exists",
			args: args[*int]{src: []*int{&src[0], &src[1], &src[2]}, e: &src[1]},
			want: true,
		},
		{
			name: "*int-not exists",
			args: args[*int]{src: []*int{&src[0], &src[1], &src[2]}, e: &src[3]},
			want: false,
		},
		{
			name: "*int-not exists 2",
			args: args[*int]{src: []*int{&src[0], &src[1], &src[2]}, e: &outOfSrc},
			want: false,
		},
	}
	for _, tt := range intPtrTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.src, tt.args.e); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}

}
