package a

func Filters[T any](src []T, fs ...func(T) bool) [][]T {
	ret := make([][]T, len(fs)+1)
	if len(fs) == 0 {
		ret[0] = src
		return ret
	}
ELEMENT_LOOP:
	for i := range src {
		isMatched := false
		for j := range fs {
			if fs[j](src[i]) {
				ret[j] = append(ret[j], src[i])
				isMatched = true
			}
		}
		if isMatched {
			continue ELEMENT_LOOP
		}
		// others; the element unmatched all filters append into the last slice.
		ret[len(fs)] = append(ret[len(fs)], src[i])
	}
	return ret
}
