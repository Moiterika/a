package a

import (
	"sort"
)

func Sort[T any](src []T, cmps ...func(t1 T, t2 T) bool) []T {
	if src == nil {
		return nil
	}
	ret := make([]T, len(src))
	copy(ret, src)

	sort.SliceStable(ret, func(i, j int) bool {
		t1 := ret[i]
		t2 := ret[j]
		for _, cmp := range cmps {
			if cmp(t1, t2) {
				return true
			}
			if cmp(t2, t1) {
				return false
			}
		}
		return false
	})
	return ret
}
