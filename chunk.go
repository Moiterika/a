package a

func Chunk[T any](src []T, size int) [][]T {
	if size <= 0 || src == nil {
		return nil
	}
	if len(src) < size {
		ret := make([][]T, 1)
		ret[0] = src
		return ret
	}

	// determine total chunk count
	cc := len(src) / size
	if len(src)%size != 0 {
		cc++
	}

	ret := make([][]T, cc)
	for n := 0; n < cc-1; n++ {
		ret[n] = src[size*n : size*(n+1)]
	}
	ret[cc-1] = src[size*(cc-1):]
	return ret
}
