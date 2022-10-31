package options

// Default
// Returns the first parameter if it's not a zero value.
// Returns the second parameter otherwise.
func Default[T comparable](value, defaultValue T) T {
	var zero T
	if value != zero {
		return value
	}
	return defaultValue
}
