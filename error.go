package au

import "fmt"

// Must is a helper function that panics if the provided error is not nil.
// It returns the value v if err is nil, otherwise it panics with a formatted error message.
// This is useful for handling errors that are not expected to occur in normal operation.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}

	return v
}
