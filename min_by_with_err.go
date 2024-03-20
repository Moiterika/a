package a

import (
	"cmp"
	"fmt"
)

func MinByWithErr[T any, S cmp.Ordered](src []T, m func(T) S) (T, error) {
	var ret T
	if m == nil {
		return ret, fmt.Errorf("m is nil: %w", ErrArg)
	}
	if len(src) == 0 {
		return ret, nil
	}
	ret = src[0]
	minVal := m(src[0])
	for i := 1; i < len(src); i++ {
		v := m(src[i])
		if minVal > v {
			ret = src[i]
			minVal = v
		}
	}
	return ret, nil
}
