package a

import "fmt"

func ToMapWithErr[K comparable, T any](src []T, k func(T) K) (map[K]T, error) {
	ret := make(map[K]T, len(src))
	if k == nil {
		return ret, fmt.Errorf("k is nil: %w", ErrArg)
	}
	for _, e := range src {
		if _, exists := ret[k(e)]; exists {
			return nil, fmt.Errorf("%v: %w", k(e), ErrDuplicateKeyValue)
		}
		ret[k(e)] = e
	}
	return ret, nil
}
