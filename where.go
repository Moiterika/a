package a

func Where[T any](src []T, w func(T) bool) []T {
	ret := make([]T, 0, len(src))
	for _, e := range src {
		if w(e) {
			ret = append(ret, e)
		}
	}
	return ret
}
