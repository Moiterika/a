package a

func GroupBy[K comparable, T any](src []T, k func(T) K) map[K][]T {
	ret := make(map[K][]T, len(src))
	for _, e := range src {
		if _, exists := ret[k(e)]; exists {
			ret[k(e)] = append(ret[k(e)], e)
		} else {
			ret[k(e)] = []T{e}
		}
	}
	return ret
}
