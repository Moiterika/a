package a

func All[T any](src []T, a func(T) bool) bool {
	if a == nil {
		return false
	}
	for _, e := range src {
		if !a(e) {
			return false
		}
	}
	return true
}
