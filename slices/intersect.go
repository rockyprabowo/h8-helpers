package slices

import (
	"github.com/rockyprabowo/h8-helpers/set"
	"golang.org/x/exp/constraints"
)

// Intersect
// Returns the intersection of two slice with a type of T without the duplicates.
func Intersect[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)

	sets := set.NewSet[T]()

	for _, v := range short {
		sets.Add(v)
	}

	for _, v := range long {
		if sets.Has(v) {
			diff = append(diff, v)
		}
	}

	return
}
