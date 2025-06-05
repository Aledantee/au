package au

// Pointer returns a pointer to the given value.
//
// Example:
//
//	n := 42
//	p := Pointer(n) // p is *int with value 42
func Pointer[T any](v T) *T {
	return &v
}

// Ptr is a shorthand alias for Pointer.
//
// Example:
//
//	n := 42
//	p := Ptr(n) // p is *int with value 42
func Ptr[T any](v T) *T {
	return &v
}

// DerefOr returns the value pointed to by p, or def if p is nil.
//
// Example:
//
//	var p *int
//	v := DerefOr(p, 42) // v is 42
func DerefOr[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}
