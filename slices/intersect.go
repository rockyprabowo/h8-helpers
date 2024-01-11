package slices

import (
	"github.com/rockyprabowo/h8-helpers/set"
	"golang.org/x/exp/constraints"
)

// Intersect
// Returns the intersection of two slice with a type of T without the duplicates.
func Intersect[T constraints.Ordered](a, b []T) []T {
	intersectSet := set.NewSet[T]()
	long, short := SliceLongShort(a, b)

	shortSet := set.NewSetFromSlice[T](short)
	longSlice := set.NewSetFromSlice[T](long).ToSlice()

	for i := range longSlice {
		if shortSet.Has(longSlice[i]) {
			intersectSet.Add(longSlice[i])
		}
	}

	return intersectSet.ToSlice()
}
