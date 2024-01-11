package slices

import (
	"github.com/rockyprabowo/h8-helpers/set"
	"golang.org/x/exp/constraints"
)

// SliceLongShort
// Returns a longer list at the first return value and shorter one on the second return value.
func SliceLongShort[T any](a, b []T) ([]T, []T) {
	if len(b) > len(a) {
		return b, a
	}
	return a, b
}

// Diff
// Returns the difference of two slice with a type of T without the duplicates.
func Diff[T constraints.Ordered](a, b []T) []T {
	var diff []T
	long, short := SliceLongShort(a, b)

	shortSet := set.NewSetFromSlice[T](short)
	longSet := set.NewSetFromSlice[T](long)

	for _, v := range longSet.ToSlice() {
		if !shortSet.Has(v) {
			diff = append(diff, v)
		}
	}

	return diff
}
