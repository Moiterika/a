package a

func Sum[T any, S number](src []T, s func(T) S) S {
	var sum S
	if len(src) == 0 || s == nil {
		return sum
	}
	for i := range src {
		sum += s(src[i])
	}
	return sum
}
