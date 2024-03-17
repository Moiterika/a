package a

import (
	"cmp"
	"sort"
)

func OrderBy[T any, S cmp.Ordered](src []T, o func(T) S) {
	sort.SliceStable(src, func(i, j int) bool {
		return o(src[i]) < o(src[j])
	})
}
