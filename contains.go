package a

func Contains[T comparable](src []T, e T) bool {
	for i := range src {
		if e == src[i] {
			return true
		}
	}
	return false
}
