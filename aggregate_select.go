package a

func AggregateSelect[T, S any](src []T, initialF func(T) S, f func(S, T) S) S {
	var ret S
	if len(src) == 0 || initialF == nil || f == nil {
		return ret
	}

	ret = initialF(src[0])

	for i := 1; i < len(src); i++ {
		ret = f(ret, src[i])
	}
	return ret
}
