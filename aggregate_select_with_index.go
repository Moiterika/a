package a

func AggregateSelectWithIndex[T, S any](src []T, initialF func(T) S, f func(S, T, int) S) S {
	var ret S
	if len(src) == 0 || initialF == nil || f == nil {
		return ret
	}

	ret = initialF(src[0])

	for i := 1; i < len(src); i++ {
		ret = f(ret, src[i], i)
	}
	return ret
}
