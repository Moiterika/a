package a

import (
	"cmp"
)

func MinBy[T any, S cmp.Ordered](src []T, m func(T) S) T {
	var ret T
	if len(src) == 0 || m == nil {
		return ret
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
	return ret
}
