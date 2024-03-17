package a

func Any[T any](src []T, a func(T) bool) bool {
	for _, e := range src {
		if a(e) {
			return true
		}
	}
	return false
}
