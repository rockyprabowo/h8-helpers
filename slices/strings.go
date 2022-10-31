package slices

// StringInSlice
// Returns true if the text is in the slice
// Returns false otherwise.
func StringInSlice(
	text string,
	slice []string,
) bool {
	for _, b := range slice {
		if b == text {
			return true
		}
	}
	return false
}
