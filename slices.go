package au

import "fmt"

// Prepend adds a value to the beginning of a slice and returns the new slice.
// It is similar to append but adds the element at the start instead of the end.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	prepended := Prepend(numbers, 0)
//	// prepended is []int{0, 1, 2, 3}
func Prepend[T any, S ~[]T](slice S, value T) S {
	return append([]T{value}, slice...)
}

// Map applies a function to each element of a slice and returns a new slice containing the results.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	doubled := Map(numbers, func(n int) int { return n * 2 })
//	// doubled is []int{2, 4, 6}
func Map[I any, O any, IS ~[]I, OS ~[]O](slice IS, fn func(I) O) OS {
	res := make(OS, len(slice))

	for i, v := range slice {
		res[i] = fn(v)
	}

	return res
}

// MapStringer converts a slice of types implementing fmt.Stringer to a slice of strings.
// It uses the String() method of each element to perform the conversion.
//
// Example:
//
//	type Person struct { Name string }
//	func (p Person) String() string { return p.Name }
//
//	people := []Person{{"Alice"}, {"Bob"}}
//	names := MapStringer(people)
//	// names is []string{"Alice", "Bob"}
func MapStringer[I fmt.Stringer, IS ~[]I](slice IS) []string {
	return Map[I, string, IS, []string](slice, func(i I) string {
		return i.String()
	})
}
