package a

func Exists[T comparable](src []T, e T) bool {
	for _, x := range src {
		if e == x {
			return true
		}
	}
	return false
}
