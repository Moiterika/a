package a

import (
	"cmp"
	"sort"
)

func OrderByDescending[T any, S cmp.Ordered](src []T, o func(T) S) []T {
	if src == nil {
		return nil
	}
	ret := make([]T, len(src))
	copy(ret, src)
	sort.SliceStable(ret, func(i, j int) bool {
		return o(ret[i]) > o(ret[j])
	})
	return ret
}
