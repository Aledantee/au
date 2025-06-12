package au

import "cmp"

// Min returns the smaller of two ordered values.
// If the values are equal, returns the first value.
//
// Example:
//
//	Min(1, 2) // 1
//	Min(2, 1) // 1
//	Min(1, 1) // 1
func Min[T cmp.Ordered](a, b T) T {
	if b < a {
		return b
	}

	return a
}

// Max returns the larger of two ordered values.
// If the values are equal, returns the first value.
//
// Example:
//
//	Max(1, 2) // 2
func Max[T cmp.Ordered](a, b T) T {
	if a < b {
		return b
	}

	return a
}
