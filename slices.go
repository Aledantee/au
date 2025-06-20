package au

import (
	"errors"
	"fmt"
	"strconv"
)

// PrependSlice adds a value to the beginning of a slice and returns the new slice.
// It is similar to append but adds the element at the start instead of the end.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	prepended := Prepend(numbers, 0)
//	// prepended is []int{0, 1, 2, 3}
func PrependSlice[T any, S ~[]T](slice S, value T) S {
	return append([]T{value}, slice...)
}

// MapSlice applies a function to each element of a slice and returns a new slice containing the results.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	doubled := Map(numbers, func(n int) int { return n * 2 })
//	// doubled is []int{2, 4, 6}
func MapSlice[I any, O any, IS ~[]I, OS ~[]O](slice IS, fn func(I) O) OS {
	res := make(OS, len(slice))

	for i, v := range slice {
		res[i] = fn(v)
	}

	return res
}

// TryMapSlice applies a function to each element of a slice and returns a new slice containing the results.
// If the function returns an error, the error is returned and the result is not added to the new slice.
//
// Example:
//
//	numbers := []int{1, 2, 3, 0}
//	doubled, err := TryMapSlice(numbers, func(n int) (int, error) {
//		if n == 0 {
//			return 0, errors.New("zero is not allowed")
//		}
//		return n * 2, nil
//	})
//	// doubled is []int{2, 4, 6}
//	// err is an error containing the error from the function call for the element 0
func TryMapSlice[I any, O any, IS ~[]I, OS ~[]O](slice IS, fn func(I) (O, error)) (OS, error) {
	var (
		res  OS
		errs []error
	)

	for _, v := range slice {
		if o, err := fn(v); err == nil {
			res = append(res, o)
		} else {
			errs = append(errs, err)
		}
	}

	return res, errors.Join(errs...)
}

// MapSliceStr converts a slice of types implementing fmt.Stringer to a slice of strings.
// It uses the String() method of each element to perform the conversion.
//
// Example:
//
//	type Person struct { Name string }
//	func (p Person) String() string { return p.Name }
//
//	people := []Person{{"Alice"}, {"Bob"}}
//	names := MapSliceStr(people)
//	// names is []string{"Alice", "Bob"}
func MapSliceStr[I fmt.Stringer, IS ~[]I](slice IS) []string {
	return MapSlice[I, string, IS, []string](slice, func(i I) string {
		return i.String()
	})
}

// MapSliceAtoi converts a slice of types implementing fmt.Stringer to a slice of ints.
// Values that cannot be converted to ints are dropped.
// It uses the String() method of each element to perform the conversion.
//
// Example:
//
//	numbers := []string{"1", "2", "3"}
func MapSliceAtoi[I fmt.Stringer, IS ~[]I](slice IS) []int {
	res, _ := TryMapSlice[I, int, IS, []int](slice, func(i I) (int, error) {
		return strconv.Atoi(i.String())
	})

	return res
}

// ChunkSlice splits a slice into smaller chunks of the specified size.
// The last chunk may be smaller than the specified size if the slice length
// is not evenly divisible by the chunk size.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	chunks := ChunkSlice(numbers, 2)
//	// chunks is [][]int{{1, 2}, {3, 4}, {5, 6}}
//
//	numbers := []int{1, 2, 3, 4, 5}
//	chunks := ChunkSlice(numbers, 3)
//	// chunks is [][]int{{1, 2, 3}, {4, 5}}
func ChunkSlice[T any, S ~[]T](slice S, size int) []S {
	if size <= 0 {
		panic("size must be greater than 0")
	}

	chunks := make([]S, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
