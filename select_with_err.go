package a

import "fmt"

func SelectWithErr[T, S any](src []T, s func(T) S) ([]S, error) {
	if s == nil {
		return nil, fmt.Errorf("s is nil: %w", ErrArg)
	}
	if src == nil {
		return nil, nil
	}
	ret := make([]S, 0, len(src))
	for i := range src {
		ret = append(ret, s(src[i]))
	}
	return ret, nil
}
