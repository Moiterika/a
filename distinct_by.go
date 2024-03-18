package a

import "fmt"

func DistinctBy[T any, K comparable](src []T, d func(T) K) ([]T, error) {
	if d == nil {
		return nil, fmt.Errorf("d is nil: %w", ErrArg)
	}
	if src == nil {
		return nil, nil
	}
	ret := make([]T, len(src))
	copy(ret, src)
	m := make(map[K]struct{})
	n := 0
	for i := range ret {
		k := d(ret[i])
		if _, exists := m[k]; !exists {
			m[k] = struct{}{}
			ret[n] = ret[i]
			n++
		}
	}
	ret = ret[:n]
	return ret, nil
}
