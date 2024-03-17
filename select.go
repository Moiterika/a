package a

func Select[T, S any](src []T, s func(T) S) []S {
	ret := make([]S, 0, len(src))
	for _, e := range src {
		ret = append(ret, s(e))
	}
	return ret
}
