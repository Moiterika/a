package a

import "cmp"

func MaxBy[T any, S cmp.Ordered](src []T, m func(T) S) T {
	var ret T
	if len(src) == 0 {
		return ret
	}
	ret = src[0]
	maxVal := m(src[0])
	for i := 1; i < len(src); i++ {
		v := m(src[i])
		if maxVal < v {
			ret = src[i]
			maxVal = v
		}
	}
	return ret
}
