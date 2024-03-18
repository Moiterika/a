package a

import "fmt"

func Any[T any](src []T, a func(T) bool) (bool, error) {
	if a == nil {
		return true, fmt.Errorf("a is nil: %w", ErrArg)
	}
	for _, e := range src {
		if a(e) {
			return true, nil
		}
	}
	return false, nil
}
