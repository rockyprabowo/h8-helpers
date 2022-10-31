package set

// Set implementation using a map of struct{}s.
//
// The map is used to store the set's elements. The struct{} is an empty type,
// which means it takes up no space. It's used as a placeholder to indicate that
// the map's key is present.
//
// The Set type is generic, which means it can be used to create sets of any
// comparable type.
type Set[T comparable] struct {
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

// NewSet creates a new set of type T
func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	return s
}

// NewSetFromSlice creates a new set of type T
func NewSetFromSlice[T comparable](slice []T) *Set[T] {
	s := &Set[T]{}
	s.sets = make(map[T]struct{})
	s.AddMany(slice...)
	return s
}
