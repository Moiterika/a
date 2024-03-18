package a

import "fmt"

func GroupBy[K comparable, T any](src []T, k func(T) K) (map[K][]T, error) {
	if k == nil {
		return nil, fmt.Errorf("k is nil: %w", ErrArg)
	}
	ret := make(map[K][]T, len(src))
	for _, e := range src {
		ret[k(e)] = []T{e}
	}
	return ret, nil
}
