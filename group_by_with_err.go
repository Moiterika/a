package a

import "fmt"

func GroupByWithErr[K comparable, T any](src []T, k func(T) K) (map[K][]T, error) {
	ret := make(map[K][]T, len(src))
	if k == nil {
		return ret, fmt.Errorf("k is nil: %w", ErrArg)
	}
	for _, e := range src {
		if _, exists := ret[k(e)]; !exists {
			ret[k(e)] = []T{e}
			continue
		}
		ret[k(e)] = append(ret[k(e)], e)
	}
	return ret, nil
}
