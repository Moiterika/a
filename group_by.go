package a

func GroupBy[K comparable, T any](src []T, k func(T) K) map[K][]T {
	if k == nil {
		return make(map[K][]T, 0)
	}
	ret := make(map[K][]T, len(src))
	for _, e := range src {
		if _, exists := ret[k(e)]; !exists {
			ret[k(e)] = []T{e}
			continue
		}
		ret[k(e)] = append(ret[k(e)], e)
	}
	return ret
}
