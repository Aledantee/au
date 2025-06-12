package au

import "github.com/maruel/natural"

// NatLte returns true if a is less than or equal to b using natural string comparison.
// Natural string comparison handles numeric parts of strings as numbers rather than lexicographically.
//
// Example:
//
//	NatLte("file1", "file2") // true
//	NatLte("file2", "file1") // false
//	NatLte("file1", "file1") // true
func NatLte(a, b string) bool {
	return a == b || natural.Less(a, b)
}

// NatGte returns true if a is greater than or equal to b using natural string comparison.
// Natural string comparison handles numeric parts of strings as numbers rather than lexicographically.
//
// Example:
//
//	NatGte("file2", "file1") // true
//	NatGte("file1", "file2") // false
//	NatGte("file1", "file1") // true
func NatGte(a, b string) bool {
	return a == b || !natural.Less(a, b)
}

// NatLt returns true if a is less than b using natural string comparison.
// Natural string comparison handles numeric parts of strings as numbers rather than lexicographically.
//
// Example:
//
//	NatLt("file1", "file2") // true
//	NatLt("file2", "file1") // false
//	NatLt("file1", "file1") // false
func NatLt(a, b string) bool {
	return a != b && natural.Less(a, b)
}

// NatGt returns true if a is greater than b using natural string comparison.
// Natural string comparison handles numeric parts of strings as numbers rather than lexicographically.
//
// Example:
//
//	NatGt("file2", "file1") // true
//	NatGt("file1", "file2") // false
//	NatGt("file1", "file1") // false
func NatGt(a, b string) bool {
	return a != b && !natural.Less(a, b)
}
