package au

import "iter"

// Prepend returns a new iterator that yields the given value first, then yields all elements from the input sequence.
// It is similar to append, but adds the element at the start instead of the end.
//
// Example:
//
//	seq := iter.SeqOf(1, 2, 3)
//	prepended := Prepend(seq, 0)
//	iter.ForEach(prepended, func(v int) { fmt.Println(v) })
//	// Output: 0 1 2 3
func Prepend[T any](seq iter.Seq[T], value T) iter.Seq[T] {
	return func(yield func(T) bool) {
		yield(value)
		seq(yield)
	}
}
