package a

func LastOrDefault[T any](src []T) T {
	if len(src) == 0 {
		var e T
		return e
	}

	return src[len(src)-1]
}
