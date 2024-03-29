package slices

// Map
// Returns a new slice mapped from the return value of function f.
func Map[T any, U any](slice []T, f func(T) U) []U {
	mapped := make([]U, len(slice))

	for i, v := range slice {
		mapped[i] = f(v)
	}

	return mapped
}

// Filter
// Returns a new slice with the elements that satisfies the boolean condition of function f.
func Filter[T any](slice []T, f func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
