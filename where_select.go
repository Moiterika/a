package a

func WhereSelect[T, S any](src []T, w func(T) bool, s func(T) S) []S {
	ret := make([]S, 0, len(src))
	for i := range src {
		if w(src[i]) {
			ret = append(ret, s(src[i]))
		}
	}
	return ret
}
