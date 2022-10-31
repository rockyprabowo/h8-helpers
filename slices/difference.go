package slices

import (
	"github.com/rockyprabowo/h8-helpers/set"
	"golang.org/x/exp/constraints"
)

// SliceLongShort
// Returns a longer list at the first return value and shorter one on the second return value.
func SliceLongShort[T any](a, b []T) (long, short []T) {
	if len(b) > len(a) {
		return b, a
	}
	return a, b
}

// Diff
// Returns the difference of two slice with a type of T without the duplicates.
func Diff[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)

	sets := set.NewSet[T]()

	for _, v := range short {
		sets.Add(v)
	}

	for _, v := range long {
		if !sets.Has(v) {
			diff = append(diff, v)
		}
	}

	return
}

// DiffWithDuplicates
// Returns the difference of two slice with a type of T including the duplicates.
func DiffWithDuplicates[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)
	longIter, shortIter := 0, 0

	for longIter < len(long) && shortIter < len(short) {
		if long[longIter] == short[longIter] {
			longIter++
			shortIter++
		} else if long[longIter] < short[shortIter] {
			diff = append(diff, long[longIter])
			longIter++
		} else {
			diff = append(diff, short[shortIter])
			shortIter++
		}
	}

	diff = append(diff, a[longIter:]...)
	diff = append(diff, b[shortIter:]...)

	return
}
