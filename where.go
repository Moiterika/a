package a

func Where[T any](src []T, w func(T) bool) []T {
	if src == nil {
		return nil
	}
	ret := make([]T, 0, len(src))
	for i := range src {
		if w(src[i]) {
			ret = append(ret, src[i])
		}
	}
	return ret
}
