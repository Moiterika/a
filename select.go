package a

func Select[T, S any](src []T, s func(T) S) []S {
	if src == nil || s == nil {
		return nil
	}
	ret := make([]S, 0, len(src))
	for i := range src {
		ret = append(ret, s(src[i]))
	}
	return ret
}
