package a

import "fmt"

func SumWithErr[T any, S number](src []T, s func(T) S) (S, error) {
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
