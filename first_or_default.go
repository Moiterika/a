package a

func FirstOrDefault[T any](src []T) T {
	if len(src) == 0 {
		var e T
		return e
	}

	return src[0]
}
