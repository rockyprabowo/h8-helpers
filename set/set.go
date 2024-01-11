package set

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Set implementation using a map of struct{}s.
//
// The map is used to store the set's elements. The struct{} is an empty type,
// which means it takes up no space. It's used as a placeholder to indicate that
// the map's key is present.
//
// The Set type is generic, which means it can be used to create sets of any
// constraint.Ordered type.
type Set[T constraints.Ordered] struct {
	sets map[T]struct{}
}

// Has checks if the set has a value.
func (s *Set[T]) Has(v T) bool {
	_, ok := s.sets[v]
	return ok
}

// Add adds a new element to the set.
func (s *Set[T]) Add(v T) {
	s.sets[v] = struct{}{}
}

// AddMany adds many new element to the set.
func (s *Set[T]) AddMany(v ...T) {
	for _, value := range v {
		s.sets[value] = struct{}{}
	}
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.sets, v)
}

// RemoveMany removes many value from the set.
func (s *Set[T]) RemoveMany(v ...T) {
	for _, value := range v {
		delete(s.sets, value)
	}
}

// Clear clears the set.
func (s *Set[T]) Clear() {
	s.sets = make(map[T]struct{})
}

// Size returns the size of the set.
func (s *Set[T]) Size() int {
	return len(s.sets)
}

// ToSlice returns a sorted slice of the set
func (s *Set[T]) ToSlice() []T {
	var slice []T
	for key := range s.sets {
		slice = append(slice, key)
	}
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// NewSet creates a new set of type T
func NewSet[T constraints.Ordered]() *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	return s
}

// NewSetFromSlice creates a new set of type T
func NewSetFromSlice[T constraints.Ordered](slice []T) *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	s.AddMany(slice...)
	return s
}
