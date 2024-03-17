package a

func WhereSelect[T, S any](src []T, w func(T) bool, s func(T) S) []S {
	ret := make([]S, 0, len(src))
	for _, e := range src {
		if w(e) {
			ret = append(ret, s(e))
		}
	}
	return ret
}
