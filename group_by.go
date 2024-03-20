package a

func GroupBy[K comparable, T any](src []T, k func(T) K) map[K][]T {
	if k == nil {
		return make(map[K][]T, 0)
	}
	ret := make(map[K][]T, len(src))
	for _, e := range src {
		ret[k(e)] = []T{e}
	}
	return ret
}
