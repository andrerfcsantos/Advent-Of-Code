package utils

// IntSet represents a set of integers
type IntSet struct {
	Set map[int]bool
}

// NewIntSet returns a newly created int set
func NewIntSet() IntSet {
	return IntSet{Set: make(map[int]bool)}
}

// Add adds an element to the int set
func (s IntSet) Add(element int) {
	s.Set[element] = true
}

// Remove removes an element from the int set
func (s IntSet) Remove(element int) {
	if s.Has(element) {
		delete(s.Set, element)
	}
}

// Has checks if an element exists on the int set
func (s IntSet) Has(element int) bool {
	_, ok := s.Set[element]
	return ok
}

// Elements returns the elements of this set as a slice of ints
func (s IntSet) Elements(element string) []int {

	var res []int
	for k := range s.Set {
		res = append(res, k)
	}
	return res
}
