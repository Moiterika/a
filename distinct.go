package a

func Distinct[T comparable](src []T) []T {
	// The sequence of src should be stable, so slices package is not adopted.
	// (Because the pre slice sort is necessary to lices.Compact or slices.CompactFunc)
	if src == nil {
		return nil
	}
	ret := make([]T, len(src))
	copy(ret, src)
	m := make(map[T]struct{})
	n := 0
	for i := range ret {
		if _, exists := m[ret[i]]; !exists {
			m[ret[i]] = struct{}{}
			ret[n] = ret[i]
			n++
		}
	}
	ret = ret[:n]
	return ret
}
