package a

import (
	"cmp"
	"fmt"
	"sort"
)

func OrderByDescending[T any, S cmp.Ordered](src []T, o func(T) S) ([]T, error) {
	if o == nil {
		return nil, fmt.Errorf("o is nil: %w", ErrArg)
	}
	if src == nil {
		return nil, nil
	}
	ret := make([]T, len(src))
	copy(ret, src)
	sort.SliceStable(ret, func(i, j int) bool {
		return o(ret[i]) > o(ret[j])
	})
	return ret, nil
}
