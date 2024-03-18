package a

import "fmt"

type number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T any, S number](src []T, s func(T) S) (S, error) {
	var sum S
	if s == nil {
		return sum, fmt.Errorf("s is nil: %w", ErrArg)
	}
	if len(src) == 0 {
		return sum, nil
	}
	for i := range src {
		sum += s(src[i])
	}
	return sum, nil
}
