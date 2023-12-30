package sliceconv

import (
	"golang.org/x/exp/constraints"
)

// Integer converts any integer-based slice to any other integer-based slice.
// No overflow checks.
func Integer[T1 constraints.Integer, T2 constraints.Integer](from []T1) []T2 {
	ret := make([]T2, len(from))
	for i := range from {
		ret[i] = T2(from[i])
	}

	return ret
}

// Float converts any float-based slice to any other float-based slice.
// No overflow checks.
func Float[T1 constraints.Float, T2 constraints.Float](from []T1) []T2 {
	ret := make([]T2, len(from))
	for i := range from {
		ret[i] = T2(from[i])
	}

	return ret
}
