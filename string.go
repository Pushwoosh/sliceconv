package sliceconv

// String converts any string-based slice to any other string-based slice.
func String[T1 ~string, T2 ~string](from []T1) []T2 {
	ret := make([]T2, len(from))
	for i := range from {
		ret[i] = T2(from[i])
	}

	return ret
}
