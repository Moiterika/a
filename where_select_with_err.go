package a

import "fmt"

func WhereSelectWithErr[T, S any](src []T, w func(T) bool, s func(T) S) ([]S, error) {
	if s == nil {
		return nil, fmt.Errorf("s is nil: %w", ErrArg)
	}
	ret := make([]S, 0, len(src))
	if w == nil {
		return ret, nil
	}
	for i := range src {
		if w(src[i]) {
			ret = append(ret, s(src[i]))
		}
	}
	return ret, nil
}
