package slices

import "golang.org/x/exp/constraints"

func Contains[T constraints.Ordered](slice []T, search T) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}
