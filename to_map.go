package a

func ToMap[K comparable, T any](src []T, k func(T) K) map[K]T {
	ret := make(map[K]T, len(src))
	for _, e := range src {
		ret[k(e)] = e
	}
	return ret
}
