package a

func Any[T any](src []T, a func(T) bool) bool {
	if a == nil {
		return false
	}
	for _, e := range src {
		if a(e) {
			return true
		}
	}
	return false
}
