package a

func Filters[T any](src []T, fs ...func(T) bool) ([][]T, []T) {
	if len(fs) == 0 {
		return nil, src
	}
	residues := make([][]T, len(fs))
	filtrate := make([]T, 0)
	for j := range fs {
		if fs[j] == nil {
			fs[j] = func(_ T) bool { return false }
		}
		residues[j] = make([]T, 0)
	}
ELEMENT_LOOP:
	for i := range src {
		isMatched := false
		for j := range fs {
			if fs[j](src[i]) {
				residues[j] = append(residues[j], src[i])
				isMatched = true
			}
		}
		if isMatched {
			continue ELEMENT_LOOP
		}
		// the element unmatched all filters append into filtrate.
		// this means such as "others", "else in if-else", "default in switch".
		filtrate = append(filtrate, src[i])
	}
	return residues, filtrate
}
