package a

type number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T any, S number](src []T, s func(T) S) S {
	var sum S
	if len(src) == 0 {
		return sum
	}
	for i := range src {
		sum += s(src[i])
	}
	return sum
}
